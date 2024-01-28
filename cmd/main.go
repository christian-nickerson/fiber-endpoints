package main

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/christian-nickerson/golang-onnx-api/internal/config"
	"github.com/christian-nickerson/golang-onnx-api/internal/logging"
	"github.com/christian-nickerson/golang-onnx-api/internal/routes"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	// Load configs
	config, err := config.LoadConfig("settings")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Setup app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Add logging
	app.Use(requestid.New())
	app.Use(logger.New(logging.LoggingConfig))

	// Add routes
	app.Use(healthcheck.New(routes.HealthCheckConfig))
	routes.AddInferenceRoutes(app)

	app.Listen(":" + config.API.Port)
}
