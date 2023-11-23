package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`         // Application environment (e.g., development, production)
	DBSource          string `mapstructure:"DB_SOURCE"`           // Database connection source
	MigrationURL      string `mapstructure:"MIGRATION_URL"`       // URL for migrations
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"` // HTTP server address
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // Add the provided path as a search location for the configuration file
	viper.SetConfigName("app") // Set the name of the configuration file to be read
	viper.SetConfigType("env") // Set the configuration file type as environment variable style

	viper.AutomaticEnv() // Automatically read environment variables

	err = viper.ReadInConfig() // Read the configuration file into viper
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config) // Unmarshal the configuration read by viper into the 'config' struct
	return
}
