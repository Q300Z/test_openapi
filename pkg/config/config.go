package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from the config file.
func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found, using default values")
		} else {
			log.Fatalf("failed to read config file: %s", err)
		}
	}
}
