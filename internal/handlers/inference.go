package handlers

import (
	"context"
	"log"

	"github.com/dmitryikh/leaves"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("fiber-endpoints-fiber")
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

// InferenceModel returns prediction from inference request
func InferenceModel(ctx context.Context, data []float64) float64 {
	ctx, span := tracer.Start(ctx, "InferenceModel")
	defer span.End()
	return model.PredictSingle(data, 0)
}
