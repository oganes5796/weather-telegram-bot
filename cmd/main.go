package main

import (
	"github.com/oganes5796/weather-bot/config"
	"github.com/oganes5796/weather-bot/internal/logger"
	"github.com/oganes5796/weather-bot/internal/telegram"
)

func main() {
	// Инициализация логера
	logger := logger.NewLogger()

	// Загрузка переменных окружения из .env
	config.LoadEnv()

	// Загрузка токена
	telegramToken := config.GetEnv("TELEGRAM_BOT_TOKEN")

	// Инициализация и запуск бота
	bot, err := telegram.NewBot(telegramToken)
	if err != nil {
		logger.Error("Error starting telegram bot", "error", err)
	}

	logger.Info("Bot starter...")
	bot.Start()
}
