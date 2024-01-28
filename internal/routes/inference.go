package routes

import (
	"github.com/christian-nickerson/golang-onnx-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func AddInferenceRoutes(router fiber.Router) {
	note := router.Group("/inference")
	note.Get("/", handlers.InferenceRequest)
}
