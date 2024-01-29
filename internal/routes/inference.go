package routes

import (
	"github.com/christian-nickerson/golang-onnx-api/internal/handlers"
	"github.com/christian-nickerson/golang-onnx-api/internal/schema"
	"github.com/gofiber/fiber/v2"
)

func AddInferenceRoutes(router fiber.Router) {
	note := router.Group("/inference")
	note.Post("/", schema.ValidateInferenceRequest, handlers.InferenceRequest)
}
