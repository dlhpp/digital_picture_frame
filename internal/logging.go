package internal

import (
	"log/slog"
	"os"
)

func SetupLogging() {
	// Set up structured logging with slog
	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: slog.LevelInfo, AddSource: false})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Info("setupLogging: structured logging initialized")
}
