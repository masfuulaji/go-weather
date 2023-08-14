package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var setkeyCmd = &cobra.Command{
	Use:   "set-key",
	Short: "Set the API key",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("api_key")
		if apiKey == "" {
			fmt.Println("API key not set")

			var inputKey string
			_, err := fmt.Scanln(&inputKey)
			if err != nil {
				fmt.Println(err)
				return
			}

			viper.Set("api_key", inputKey)
			if err := viper.WriteConfig(); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("API key set")

		}else{
            fmt.Println("API key set")
		}
	},
}

func init() {
	rootCmd.AddCommand(setkeyCmd)
}
