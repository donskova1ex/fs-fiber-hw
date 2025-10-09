package logger

import (
	"fiber-hw/config"
	"log/slog"
	"os"
)

func NewLogger(loggerCfg *config.LogConfig) *slog.Logger {
	level := parseLogLevel(loggerCfg.Level)
	var loggerHandler slog.Handler
	
	switch loggerCfg.Format {
	case "json":
		loggerHandler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	default:
		loggerHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}
	return slog.New(loggerHandler)
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
