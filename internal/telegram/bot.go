package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/oganes5796/weather-bot/internal/weather"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{api: api}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			city := update.Message.Text
			weatherInfo, err := weather.GetForecastWeather(city)
			if err != nil {
				weatherInfo = "Не удалось получить погоду. Проверьте название города."
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, weatherInfo)
			b.api.Send(msg)
		}
	}
}
