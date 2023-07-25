package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

// config func to get config var from a .env file
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("error loading .env file")
	}
	return os.Getenv(key)
}