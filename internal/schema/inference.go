package schema

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type InferenceRequest struct {
	Data []float64 `json:"data" validate:"required,len=20"`
}

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func ValidateInferenceRequest(c *fiber.Ctx) error {
	var errors []*IError
	body := new(InferenceRequest)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
