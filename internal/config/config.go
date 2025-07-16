package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	AIProvider      string  `mapstructure:"AI_PROVIDER"`
	AIAPIURL        string  `mapstructure:"AI_API_URL"`
	AIToken         string  `mapstructure:"AI_TOKEN"`
	AIModel         string  `mapstructure:"AI_MODEL"`
	OpenAIAPIKey    string  `mapstructure:"OPENAI_API_KEY"`
	OpenAIToken     string  `mapstructure:"OPENAI_TOKEN"`
	OpenAIBaseURL   string  `mapstructure:"OPENAI_BASE_URL"`
	OpenAIModel     string  `mapstructure:"OPENAI_MODEL"`
	OpenAITemperature float32 `mapstructure:"OPENAI_TEMPERATURE"`
	OllamaURL       string  `mapstructure:"OLLAMA_URL"`
	HTTPPort        string  `mapstructure:"HTTP_PORT"`
	LogLevel        string  `mapstructure:"LOG_LEVEL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("config file not found, using environment variables")
		} else {
			// Config file was found but another error was produced
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
