package main

import (
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.opentelemetry.io/otel"

	"github.com/christian-nickerson/fiber-endpoints/internal/config"
	"github.com/christian-nickerson/fiber-endpoints/internal/handlers"
	"github.com/christian-nickerson/fiber-endpoints/internal/logging"
	"github.com/christian-nickerson/fiber-endpoints/internal/routes"
	"github.com/christian-nickerson/fiber-endpoints/internal/tracing"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	// load configs
	config, err := config.LoadConfig("settings")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// load model into memory
	handlers.LoadModel(config.Model.File)

	// set up otel tracing
	tracerProvider, err := tracing.InitTracerProvider(config.OTEL.Host, config.Fiber.Name)
	if err != nil {
		log.Fatalf("Failed to initialize tracer provider: %v", err)
	}
	otel.SetTracerProvider(tracerProvider)

	// set up fiber app
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// add tracing
	app.Use(otelfiber.Middleware())

	// add logging
	app.Use(requestid.New())
	app.Use(logger.New(logging.LoggingConfig))

	// add routes
	app.Use(healthcheck.New(routes.HealthCheckConfig))
	routes.AddInferenceRoutes(app)

	// run
	app.Listen(":" + config.Fiber.Port)
}
