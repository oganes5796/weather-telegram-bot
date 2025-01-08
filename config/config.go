package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/oganes5796/weather-bot/internal/logger"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logger.NewLogger().Error("No .env file found", "error", err)
		return
	}
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		logger.NewLogger().Error("Missing environment variable", "key", key)
		os.Exit(1)
	}
	return value
}
