package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetGeolocation(ip string, apiKey string) (*LocationData, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location LocationData
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}
	return &location, nil
}

func GetWeatherData(cityQuery string, apiKey string) (*WeatherInfo, error) {
	if cityQuery == "" {
		cityQuery = "auto"
	}
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, cityQuery)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherData WeatherInfo
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}
	return &weatherData, nil
}

func GenerateGreeting(name, location string, temperature float32) string {
	return fmt.Sprintf("Hello, %s! The temperature is %.1f degrees Celsius in %s", name, temperature, location)
}
