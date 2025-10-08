package logger

import (
	"fiber-hw/config"
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(loggerCfg *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(loggerCfg.Level))
	var logger zerolog.Logger
	switch loggerCfg.Format {
	case "json":
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	default:
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	}
	return &logger
}
