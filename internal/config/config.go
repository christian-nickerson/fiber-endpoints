package config

import (
	"fmt"
	"os"

	"github.com/golobby/config"
	"github.com/golobby/config/v3/pkg/feeder"
)

type settings struct {
	modelFIleName string
}

func getSettings() settings {
	settings := settings{}
	tomlFeeder := feeder.Toml{Path: "settings.toml"}
	c := config.New()
	c.AddFeeder(tomlFeeder)
	c.AddStruct(&settings)
	err := c.Feed()
	if err != nil {
		fmt.Println("error loading settings.toml file:", err)
		os.Exit(1)
	}
	return settings
}
