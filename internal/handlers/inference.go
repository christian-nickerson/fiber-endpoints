package handlers

import (
	"log"

	"github.com/dmitryikh/leaves"
	"github.com/gofiber/fiber/v2"
)

var model *leaves.Ensemble
var err error

// LoadModel loads the local model file
func LoadModel() {
	useTransformation := true
	model, err = leaves.LGEnsembleFromFile("./modelling/model.txt", useTransformation)
	if err != nil {
		log.Fatal(err)
	}
}

// InferenceRequest handles an inference request
func InferenceRequest(c *fiber.Ctx) error {
	fvals := []float64{0.20543991, -0.97049844, -0.81403429, -0.23842689, -0.60704084, -0.48541492, 0.53113006, 2.01834338, -0.90745243, -1.85859731, -1.02334791, -0.6877744, 0.60984819, -0.70630121, -1.29161497, 1.32385441, 1.42150747, 1.26567231, 2.56569098, -0.11154792}
	p := model.PredictSingle(fvals, 0)
	return c.JSON(fiber.Map{"prediction": p})
}
