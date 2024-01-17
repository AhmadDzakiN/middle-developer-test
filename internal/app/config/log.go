package config

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func NewLogger() zerolog.Logger {
	log.Logger = log.With().Caller().Logger()
	zerolog.TimeFieldFormat = time.RFC3339

	if os.Getenv("LOG_PRETTY") == "true" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, FormatTimestamp: func(i interface{}) string { return time.Now().Format(time.RFC3339) }})
	}

	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return log.Logger
}
