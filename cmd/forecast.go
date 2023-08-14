package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get the weather for a city",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if forecastCity == "" {
			forecastCity = "malang"
		}
	},
	Run: getForecast,
}

func init() {
	rootCmd.AddCommand(forecastCmd)
	forecastCmd.Flags().StringVarP(&forecastCity, "city", "c", "", "City name")
}

var forecastCity string

type ForecastResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		}
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				DailyChanceOfRain int `json:"daily_chance_of_rain"`
				Condition         struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
			Hour []struct {
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				}
				ChanceOfRain int `json:"chance_of_rain"`
			}
		} `json:"forecastday"`
	} `json:"forecast"`
}

func getForecast(cmd *cobra.Command, args []string) {
	apiKey := viper.GetString("api_key")
    if apiKey == "" {
        fmt.Println("API key not set")
        fmt.Println("Please set your API key using `weather-check set-key`")
        return
    }
    weatherAPIURL := viper.GetString("weather_url")
	url := fmt.Sprintf("%s/%s?key=%s&q=%s&days=1", weatherAPIURL, "forecast.json", apiKey, forecastCity)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get weather : ", err)
		return
	}
	defer response.Body.Close()

	var forecastResponse ForecastResponse
	if err := json.NewDecoder(response.Body).Decode(&forecastResponse); err != nil {
		fmt.Println("Failed to decode weather : ", err)
		return
	}

	fmt.Printf("Weather in %s:\n", forecastCity)
	fmt.Printf("Weather Now: %s\n", forecastResponse.Current.Condition.Text)
	fmt.Printf("Temperature Now: %f\n", forecastResponse.Current.TempC)
	fmt.Printf("======================================\n")
	for i := 0; i < len(forecastResponse.Forecast.Forecastday); i++ {
		parsedDate, err := time.Parse("2006-01-02", forecastResponse.Forecast.Forecastday[i].Date)
        if err != nil {
            fmt.Println(err)
        }
		fmt.Printf("Date %s:\n", parsedDate.Format("02-01-2006"))
		fmt.Printf("Description: %s\n", forecastResponse.Forecast.Forecastday[i].Day.Condition.Text)
		fmt.Printf("Chance of rain: %d\n", forecastResponse.Forecast.Forecastday[i].Day.DailyChanceOfRain)
		for j := 0; j < len(forecastResponse.Forecast.Forecastday[i].Hour); j += 4 {
			parsedTime, err := time.Parse("2006-01-02 15:04", forecastResponse.Forecast.Forecastday[i].Hour[j].Time)
			if err != nil {
				fmt.Println(err)
			}
            forecast := fmt.Sprintf(
                "%s - %.0fC, %d%%, %s", 
                parsedTime.Format("15:04"),
                forecastResponse.Forecast.Forecastday[i].Hour[j].TempC,
                forecastResponse.Forecast.Forecastday[i].Hour[j].ChanceOfRain,
                forecastResponse.Forecast.Forecastday[i].Hour[j].Condition.Text)
            if forecastResponse.Forecast.Forecastday[i].Hour[j].ChanceOfRain > 40 {
                color.Blue(forecast)
                continue
            }
            if forecastResponse.Forecast.Forecastday[i].Hour[j].Condition.Text == "Sunny" {
                color.Yellow(forecast)
                continue
            }
            fmt.Println(forecast)
		}
		fmt.Printf("======================================\n")
	}
}
