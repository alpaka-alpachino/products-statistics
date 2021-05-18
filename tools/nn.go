package tools

import (
	"github.com/dathoangnd/gonet"
	"log"
)

// CreateModel creates neural network model and saves with asked name
func CreateModel(inputs int, trainingData [][][]float64, model string) error {
	nn := gonet.New(inputs, []int{20, 5}, 1, false)

	nn.Train(trainingData, 100000, 0.4, 0.2, true)

	err := nn.Save(model)
	if err != nil {
		return err
	}

	return nil
}

// PredictForModel makes prediction using asked neural network model
func PredictForModel(inputs []float64, model string) []float64 {
	nn2, err := gonet.Load(model)
	if err != nil {
		log.Fatal("Load model failed.")
	}

	return nn2.Predict(inputs)
}
