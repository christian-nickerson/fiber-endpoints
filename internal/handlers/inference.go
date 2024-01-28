package handlers

import "github.com/gofiber/fiber/v2"

// InferenceRequest handles an inference request
func InferenceRequest(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
