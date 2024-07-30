package routes

import (
	"github.com/christian-nickerson/fiber-endpoints/internal/handlers"
	"github.com/christian-nickerson/fiber-endpoints/internal/schema"
	"github.com/gofiber/fiber/v2"
)

func AddInferenceRoutes(router fiber.Router) {
	inference := router.Group("/inference")

	// inference endpoint
	inference.Post("", schema.ValidateInferenceRequest, func(context *fiber.Ctx) error {
		body := new(schema.InferenceRequest)
		context.BodyParser(&body)
		p := handlers.TracedInference(context.UserContext(), body.Data)
		return context.Status(fiber.StatusOK).JSON(fiber.Map{"prediction": p})
	})
}
