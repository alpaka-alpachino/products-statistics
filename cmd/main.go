package main

import (
	"fmt"
	"github.com/web-site/statistics"
	"github.com/web-site/storage"
	"github.com/web-site/tools"
	"log"
)

func main() {
	fmt.Println(statistics.GetHotRises())

	targets, targetRowsLen := statistics.GetSideStatForNN(storage.KeyPoint)
	stats, sideRowsLen := statistics.GetSideStatForNN(storage.InterestPointsNN)
	if targetRowsLen != sideRowsLen {
		log.Println("input rows quantity must be equal for target and side products")
	}

	trainingData := statistics.GetTrainingData(targets,stats, targetRowsLen)

	tools.CreateModel(8, trainingData, "model.json")

	 prediction :=tools.PredictForModel(
		 []float64{0.10170, 0.052899999999999996, 0.0895, 0.0756, 0.039130000000000005, 0.03716, 0.03163, 0.03174},
		 "model.json")

	fmt.Println(prediction)

}