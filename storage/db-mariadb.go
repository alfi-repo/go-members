package storage

import (
	"database/sql"
	"go-members/config"
	"time"

	"github.com/rs/zerolog/log"
)

func NewMariaDB(cfg config.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("database connection failed")
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(cfg.DB.MaxIdleSecond))
	db.SetMaxOpenConns(cfg.DB.MaxOpenPool)
	db.SetMaxIdleConns(cfg.DB.MaxIdlePool)
	return db
}
