package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func healthCheckProbe(*fiber.Ctx) bool { return true }

func readinessProbe(*fiber.Ctx) bool { return true }

var HealthCheckConfig = healthcheck.Config{
	LivenessProbe:     healthCheckProbe,
	ReadinessProbe:    readinessProbe,
	LivenessEndpoint:  "/health",
	ReadinessEndpoint: "/ready",
}
