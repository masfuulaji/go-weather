package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
