package main

import (
	"log"

	"github.com/christian-nickerson/golang-onnx-api/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configs
	config, err := config.LoadConfig("settings")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	app := fiber.New()

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP! for now... or will it...?")
		return err
	})

	app.Listen(":" + config.API.Port)
}
