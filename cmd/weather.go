package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	weatherCmd.Flags().StringVarP(&city, "city", "c", "", "City name")
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

func getWeather(cmd *cobra.Command, args []string) {
    apiKey := viper.GetString("api_key") 
    if apiKey == "" {
        fmt.Println("API key not set")
        fmt.Println("Please set your API key using `weather-check set-key`")
        return
    }
    weatherAPIURL := viper.GetString("weather_url")

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
