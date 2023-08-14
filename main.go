package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/masfuulaji/go-weather/cmd"
	"github.com/spf13/viper"
)



func main() {
    initConfig()
    cmd.Execute()

}

func initConfig() {
	viper.SetConfigType("yaml") // Set the configuration type

	// Set the config file name and path
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".config", "weather-check")
	configFile := "config.yaml"

    configFileDir := filepath.Join(configDir, configFile)

	// Create the config directory if it doesn't exist
    _, err := os.Stat(configFileDir);
    notExist := os.IsNotExist(err)
    if notExist {
        _, err = os.Create(configFileDir)
        if err != nil {
            fmt.Println("Error creating config directory:", err)
        }
    }

	viper.SetConfigFile(filepath.Join(configDir, configFile))

	// Read the config file
    err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error reading config:", err)
			os.Exit(1)
		}
	}

	// Set default values
	viper.SetDefault("api_key", "")
	viper.SetDefault("weather_url", "http://api.weatherapi.com/v1")
}
