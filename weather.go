package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
    Use:   "weather",
    Short: "Get the weather for a city",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        if city == "" {
            city = "malang"
        }
    },
    Run: getWeather,
}

func init() {
    rootCmd.AddCommand(weatherCmd)
}

var city string

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Humidity int     `json:"humidity"`
        Condition struct {
            Text string `json:"text"`
        }
	} `json:"current"`
}

func init() {
	weatherCmd.Flags().StringVarP(&city, "city", "c", "", "City name")
}

func getWeather(cmd *cobra.Command, args []string) {
    apiKey := API_KEY 
	url := fmt.Sprintf("%s/%s?key=%s&q=%s", weatherAPIURL,"current.json", apiKey, city)

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
    fmt.Printf("Description: %s\n", weatherResponse.Current.Condition.Text)
    fmt.Printf("Temperature: %f\n", weatherResponse.Current.TempC)
    fmt.Printf("Humidity: %d\n", weatherResponse.Current.Humidity)
}
