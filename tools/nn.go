package tools

import (
	"github.com/dathoangnd/gonet"
	"log"
)

func CreateModel(inputs int, trainingData [][][]float64, model string) {
	// Create a neural network
	// 2 nodes in the input layer
	// 2 hidden layers with 4 nodes each
	// 1 node in the output layer
	// The problem is classification, not regression
	nn := gonet.New(inputs, []int{20, 5}, 1, false)

	// Train the network
	// Run for 3000 epochs
	// The learning rate is 0.4 and the momentum factor is 0.2
	// Enable debug mode to log learning error every 1000 iterations
	nn.Train(trainingData, 100000, 0.4, 0.2, true)

	// Save the model
	nn.Save(model)
}

func PredictForModel(inputs []float64, model string) []float64 {
	// Load the model
	nn2, err := gonet.Load(model)
	if err != nil {
		log.Fatal("Load model failed.")
	}
	//testInput := []float64{1, 1, 0}
	return nn2.Predict(inputs)
	//fmt.Printf("%f, %f, %f => %f\n", testInput[0], testInput[1], testInput[2], nn2.Predict(testInput)[0])
}
