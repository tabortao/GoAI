package utils

import (
	"log/slog"
	"os"
	"strings"
)

// NewLogger creates and returns a new slog.Logger instance based on the provided log level.
func NewLogger(levelStr string) *slog.Logger {
	var level slog.Level

	switch strings.ToUpper(levelStr) {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	return slog.New(handler)
}