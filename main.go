package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	weatherAPIURL = "http://api.openweathermap.org/data/2.5/weather"
)

var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get the weather for a city",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if city == "" {
			city = "Malang"
		}
	},
	Run: getWeather,
}

var city string

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
}

func init() {
	rootCmd.Flags().StringVarP(&city, "city", "c", "", "City name")
}

func getWeather(cmd *cobra.Command, args []string) {
    apiKey := API_KEY 
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", weatherAPIURL, city, apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get weather : ", err)
		return
	}
	defer response.Body.Close()

    var weatherResponse WeatherResponse
    if err := json.NewDecoder(response.Body).Decode(&weatherResponse); err != nil {
        fmt.Println("Failed to decode weather : ", err)
        return
    }

    fmt.Printf("Weather in %s:\n", city)
    fmt.Printf("Description: %s\n", weatherResponse.Weather[0].Description)
    fmt.Printf("Temperature: %f\n", weatherResponse.Main.Temp)
    fmt.Printf("Humidity: %d\n", weatherResponse.Main.Humidity)
}

func main() {
    err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
