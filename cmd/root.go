package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "weather-cli",
	Short: "CLI tool for weather",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        if city == "" {
            city = "malang"
        }
    },
    Run: getWeather,
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
