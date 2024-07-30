package handlers

import (
	"context"
	"log"

	"github.com/dmitryikh/leaves"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/christian-nickerson/fiber-endpoints")
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

// TracedInference returns prediction from inference request with tracing
func TracedInference(context context.Context, data []float64) float64 {
	context, span := tracer.Start(context, "InferenceModel")
	defer span.End()
	return model.PredictSingle(data, 0)
}
