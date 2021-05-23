package main

import (
	"fmt"
	"github.com/web-site/statistics"
	"github.com/web-site/storage"
	"github.com/web-site/tools"
	"log"
)

var (
	hotRises              []string
	regionsForSelected    map[string]string
	selected              string
	predictionForSelected string
	err                   error
)

func main() {
	hotRises, err = statistics.GetChanges(storage.All, statistics.HotRises)
	if err != nil {
		log.Fatal(err)
	}

	selected = storage.KeyPoint[0]
	regionsForSelected = statistics.GetRegionsForProducts(storage.KeyPoint)

	targets, targetRowsLen, err := statistics.GetStatsForNN(storage.KeyPoint)
	if err != nil {
		log.Fatal(err)
	}
	stats, sideRowsLen, err := statistics.GetStatsForNN(storage.InterestPointsNN)
	if err != nil {
		log.Fatal(err)
	}

	if targetRowsLen != sideRowsLen {
		log.Fatal("input rows quantity must be equal for target and side products")
	}

	trainingData := statistics.GetTrainingData(targets, stats, targetRowsLen)

	err = tools.CreateModel(8, trainingData, "model_for_chicken_meat.json")
	if err != nil {
		log.Fatal(err)
	}

	denormalizedPrediction := tools.PredictForModel(
		[]float64{0.10415, 0.05975, 0.07890000000000001, 0.0713, 0.038130000000000004, 0.0375, 0.027899999999999998, 0.02889},
		"model_for_chicken_meat.json")[0] * 1000
	predictionForSelected = fmt.Sprintf("%.2f", denormalizedPrediction)

	handleFunc()
}
