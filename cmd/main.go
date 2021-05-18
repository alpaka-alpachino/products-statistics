package main

import (
	"fmt"
	"github.com/web-site/statistics"
	"github.com/web-site/storage"
	"github.com/web-site/tools"
	"log"
)

func main() {
	fmt.Println(statistics.GetChanges(storage.All, statistics.HotCheapening))

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

	prediction := tools.PredictForModel(
		[]float64{0.10170, 0.052899999999999996, 0.0895, 0.0756, 0.039130000000000005, 0.03716, 0.03163, 0.03174},
		"model_for_chicken_meat.json")

	fmt.Println(prediction)

}
