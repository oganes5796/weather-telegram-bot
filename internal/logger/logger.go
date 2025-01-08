package logger

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	// Создаем JSON-хендлер
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Устанавливаем как глобальный логгер
	slog.SetDefault(logger)
	return logger
}
