package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	// viper.SetConfigFile("yaml") Why it doesn't working

	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(err)
	}
}
