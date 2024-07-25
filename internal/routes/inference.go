package routes

import (
	"github.com/christian-nickerson/golang-onnx-api/internal/handlers"
	"github.com/christian-nickerson/golang-onnx-api/internal/schema"
	"github.com/gofiber/fiber/v2"
)

func AddInferenceRoutes(router fiber.Router) {
	inference := router.Group("/inference")

	// inference endpoint
	inference.Post("/", schema.ValidateInferenceRequest, func(c *fiber.Ctx) error {
		body := new(schema.InferenceRequest)
		c.BodyParser(&body)
		p := handlers.InferenceModel(c.UserContext(), body.Data)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"prediction": p})
	})
}
