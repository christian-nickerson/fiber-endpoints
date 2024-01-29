package handlers

import (
	"log"

	"github.com/christian-nickerson/golang-onnx-api/internal/schema"
	"github.com/dmitryikh/leaves"
	"github.com/gofiber/fiber/v2"
)

var model *leaves.Ensemble
var err error

// LoadModel loads the local model file
func LoadModel(fileName string) {
	useTransformation := true
	model, err = leaves.LGEnsembleFromFile(fileName, useTransformation)
	if err != nil {
		log.Fatal(err)
	}
}

// InferenceRequest returns prediction from inference request
func InferenceRequest(c *fiber.Ctx) error {
	body := new(schema.InferenceRequest)
	c.BodyParser(&body)
	p := model.PredictSingle(body.Data, 0)
	return c.JSON(fiber.Map{"prediction": p})
}
