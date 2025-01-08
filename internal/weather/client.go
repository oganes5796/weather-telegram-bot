package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oganes5796/weather-bot/config"
	"github.com/oganes5796/weather-bot/internal/logger"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func GetWeather(city string) (string, error) {
	apiKey := config.GetEnv("WEATHER_API_KEY")
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&lang=ru&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get weather data: %s", resp.Status)
	}

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return "", err
	}

	weather := fmt.Sprintf("Погода в %s: %s, %.1f°C", weatherData.Name, weatherData.Weather[0].Description, weatherData.Main.Temp)
	logger.NewLogger().Info(weather)
	return weather, nil
}

func GetForecastWeather(city string) (string, error) {
	apiKey := config.GetEnv("WEATHER_API_KEY")
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=metric&lang=ru", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("не удалось получить данные о погоде, код ответа: %d", resp.StatusCode)
	}

	var forecastData struct {
		List []struct {
			DtTxt string `json:"dt_txt"`
			Main  struct {
				Temp float64 `json:"temp"`
			} `json:"main"`
			Weather []struct {
				Description string `json:"description"`
			} `json:"weather"`
		} `json:"list"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&forecastData); err != nil {
		return "", err
	}

	// Формируем строку с прогнозом на ближайшие 3 дня
	result := "Прогноз погоды:\n"
	for i, entry := range forecastData.List {
		if i%8 == 0 { // Берём прогноз каждые 8 записей (раз в день)
			result += fmt.Sprintf("%s: %.1f°C, %s\n", entry.DtTxt, entry.Main.Temp, entry.Weather[0].Description)
		}
	}

	logger.NewLogger().Info(result)
	return result, nil
}
