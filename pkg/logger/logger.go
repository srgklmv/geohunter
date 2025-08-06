package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Init() {
	// TODO: Add options when needed.
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}
