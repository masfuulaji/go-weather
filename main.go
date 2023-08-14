package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	weatherAPIURL = "http://api.weatherapi.com/v1"
)

var rootCmd = &cobra.Command{
	Use:   "weather-cli",
	Short: "CLI tool for weather",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}


func main() {
    Execute()
}
