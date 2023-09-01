package main

import (
	"context"
	"errors"
	"go-members/config"
	"go-members/internal/member"
	"go-members/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.NewConfig()
	db := storage.NewMariaDB(cfg)

	// Apps.
	memberRepoMariadb := member.NewRepoMariaDB(db)
	memberService := member.NewService(memberRepoMariadb)
	memberHandler := member.NewHandler(memberService)

	// Router init.
	httpLogger := httplog.NewLogger("members-http-logger", httplog.Options{
		JSON:    true,
		Concise: true,
	})
	routers := chi.NewRouter()
	routers.Use(
		httplog.RequestLogger(httpLogger),
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Heartbeat("/ping"),
		middleware.RealIP,
	)

	// Apps router.
	routers.Mount("/v1/members", memberHandler.Router())

	// HTTP server.
	server := &http.Server{
		Addr:              "127.0.0.1:3000",
		Handler:           routers,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal().Msg("graceful shutdown timed out.. forcing exit.")
			}
		}()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatal().Err(err).Msg("graceful shutdown error while shutdown")
		}
		serverStopCtx()
	}()

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("error server")
	}

	<-serverCtx.Done()
}
