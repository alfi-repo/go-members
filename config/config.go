package config

import (
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type appConfig struct {
	Name    string `env:"APP_NAME"`
	Address string `env:"APP_ADDRESS"`
	Debug   bool   `env:"APP_DEBUG"`
}

type dBConfig struct {
	DSN           string `env:"DB_DSN"`
	MaxOpenPool   int    `env:"DB_MAX_OPEN_POOL" envDefault:"10"`
	MaxIdlePool   int    `env:"DB_MAX_IDLE_POOL" envDefault:"10"`
	MaxIdleSecond int    `env:"DB_MAX_IDLE_SECOND" envDefault:"180"`
}

type Config struct {
	App appConfig
	DB  dBConfig
}

func NewConfig() Config {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"

	if os.Getenv("APP_ADDRESS") == "" {
		log.Info().Msg("System env not found. Load .env file")
		if err := godotenv.Load(); err != nil {
			log.Fatal().Err(err).Msg("Failed to load .env")
		}
	} else {
		log.Info().Msg("System env found")
	}

	cfg := Config{}
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(&cfg, opts); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}
	return cfg
}
