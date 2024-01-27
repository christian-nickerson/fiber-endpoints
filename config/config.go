package config

import (
	"strings"

	"github.com/spf13/viper"
)

type ApiConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	API ApiConfig `mapstructure:"api"`
}

// loadConfig reads configuration variables from toml or environment variables
func LoadConfig(name string) (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("API")

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
