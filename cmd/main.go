package main

import (
	"fmt"
	"github.com/web-site/statistics"
	"github.com/web-site/storage"
	"github.com/web-site/tools"
	"time"
)

func main() {
	//get general statistics for hot prices
	//var hotRises []string
	//var others []string
	//var interest = storage.All
	//for v := range interest {
	//	_, _, firstPrice, _, lastPrice, _, _, avg:= statistics.GetStatForProduct(interest[v])
	//	status,change := statistics.SimpleTrendCheck(firstPrice,avg,lastPrice)
	//	if status == "Hot Rise"{
	//		hotRises = append(hotRises, fmt.Sprintf(`%s:  %v`,interest[v],change))
	//	}else {
	//		others =  append(others, fmt.Sprintf(`%s:  %v`,interest[v],change))
	//	}
	//}

	targets, rawsLen := statistics.GetSideStatForNN(storage.KeyPoint)
	var targetPrices []float64
	var targetDates []time.Time
	for i := 0; i < rawsLen; i++ {
		for _, v := range targets {
			targetPrices = append(targetPrices, v.Raws[i].Prices)
			targetDates = append(targetDates, v.Raws[i].Dates)
		}
	}

	normtargetDates := statistics.DateNormalizer(targetDates)
	targetPrices = statistics.Normalizer(targetPrices)
	stats, rawsLen := statistics.GetSideStatForNN(storage.InterestPointsNN)

	var normalizedSideProds [][]float64

	for i := 0; i < rawsLen; i++{
		var normalizedSidePrices []float64

		for _, v := range stats {
			normalizedSidePrices = append(normalizedSidePrices, v.Raws[i].Prices)
		}
		normalizedSidePrices = append(normalizedSidePrices, normtargetDates[i])
		normalizedSideProds = append(normalizedSideProds, statistics.Normalizer(normalizedSidePrices))

	}
	fmt.Println("FUUCK \n",normalizedSideProds)

	var trainingData [][][]float64

	for i, v := range normalizedSideProds {
		var test [][]float64
		test = append(test, v)
		test = append(test, []float64{targetPrices[i]})
		trainingData = append(trainingData, test)
		//fmt.Println(trainingData[i], v)
	}

	fmt.Println("fuck \n",trainingData)
	tools.CreateModel(5, trainingData, "model.json")
	fmt.Println(
		tools.PredictForModel([]float64{1, 0.5700398525084732, 0.7554843755819584, 0.6818875935789042, 0}, "model.json"), "expected",0.7529691211401428,
	)
	fmt.Println(
		tools.PredictForModel([]float64{1, 0.5956765965283852, 0.8725079423546204, 0.7108982918182236, 0}, "model.json"), "expected",0.29334916864608057,
	)
	fmt.Println(
		tools.PredictForModel([]float64{1, 0.5206180494045861, 0.8611357149886822, 0.7017025883279204, 0}, "model.json"), "expected",0.45130641330166227,
	)
}
