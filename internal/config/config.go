package config

import (
	"strings"

	"github.com/spf13/viper"
)

type ApiConfig struct {
	Port string `mapstructure:"port"`
}

type ModelConfig struct {
	File string `mapstructure:"file"`
}

type Config struct {
	Fiber ApiConfig   `mapstructure:"fiber"`
	Model ModelConfig `mapstructure:"model"`
}

// loadConfig reads configuration variables from toml or environment variables
func LoadConfig(name string) (config Config, err error) {
	viper.AddConfigPath("./internal/config")
	viper.AddConfigPath(".")

	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("FIBER")

	viper.SetConfigName(name)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
