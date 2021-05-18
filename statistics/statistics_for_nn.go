package statistics

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/web-site/storage"
)

type SideNNStatsModel struct {
	Raws []Raws
}

type Raws struct {
	Dates  time.Time
	Prices float64
}

// GetTrainingData
func GetTrainingData(targets, stats []SideNNStatsModel, rawsLen int) [][][]float64{
	var targetPrices []float64
	//var targetDates []time.Time
	for i := 0; i < rawsLen; i++ {
		for _, v := range targets {
			targetPrices = append(targetPrices, v.Raws[i].Prices)
			//targetDates = append(targetDates, v.Raws[i].Dates)
		}
	}
	//statistics.DateNormalizer(targetDates)

	targetPrices = Normalizer(targetPrices)

	for _, stat := range stats {
		prices := make([]float64, 0, len(stat.Raws))
		for _, raw := range stat.Raws {
			prices = append(prices, raw.Prices)
		}

		normalizedPrices := Normalizer(prices)

		for i, v := range normalizedPrices {
			stat.Raws[i].Prices = v
		}
	}

	var normalizedSideProds [][]float64
	for i := 0; i < rawsLen; i++ {
		var normalizedSidePrices []float64

		for _, v := range stats {
			normalizedSidePrices = append(normalizedSidePrices, v.Raws[i].Prices)
		}
		normalizedSideProds = append(normalizedSideProds, normalizedSidePrices)

	}

	var trainingData [][][]float64
	targetPrices = targetPrices[1:]
	normalizedSideProds = normalizedSideProds[:len(normalizedSideProds) -1]

	for i, v := range normalizedSideProds {
		var test [][]float64
		test = append(test, v)
		test = append(test, []float64{targetPrices[i]})
		trainingData = append(trainingData, test)
	}

	return trainingData
}

//at least 2 products
func GetSideStatForNN(prods []string) ([]SideNNStatsModel, int) {
	var rawsLen int
	s := fmt.Sprintf(sourceFormat, os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DB"))
	database, err := sql.Open("postgres", s)
	if err != nil {
		panic(err)
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}(database)

	var prodSymbols string = " $1"
	for i := 2; i <= len(prods); i++ {
		prodSymbols = fmt.Sprintf("%s OR product = $%d", prodSymbols, i)
	}

	var prodValues []interface{}
	for v := range prods {
		prodValues = append(prodValues, prods[v])
	}

	set := make([]SideNNStatsModel, 0, 0)
	for p := range prodValues {
		res, err := database.Query("SELECT date, price FROM prices WHERE product = $1 ORDER BY date ASC", prodValues[p])
		if err != nil {
			panic(err)
		}

		var raws []Raws
		for res.Next() {
			var model = new(storage.Model)
			err = res.Scan(&model.Date, &model.Price)
			if err != nil {
				panic(err)
			}
			res.Err()
			storage.Statistics = append(storage.Statistics, *model)
			raws = append(raws, Raws{
				Dates:  model.Date,
				Prices: model.Price,
			})

			rawsLen = len(raws) - 1 //todo fix in scraper
		}
		//set[prods[p]] = SideNNStatsModel{Raws: raws[:len(raws)-1]}
		set = append(set, SideNNStatsModel{Raws: raws[:len(raws)-1]})

	}

	return set, rawsLen
}

func RawNormilizer(raws []Raws) ([]float64, []float64) {
	var dates []time.Time
	var prices []float64

	for _, p := range raws {
		dates = append(dates, p.Dates)
		prices = append(prices, p.Prices)
	}

	return DateNormalizer(dates), Normalizer(prices)
}

func GetNormalizedPrices(raws []Raws) []float64 {
	var prices []float64

	for _, p := range raws {
		prices = append(prices, p.Prices)
	}

	return Normalizer(prices)
}

func DateNormalizer(dates []time.Time) []float64 {
	var normalizedDates []float64
	for i := 0; i <= len(dates); i++ {
		normalizedDates = append(normalizedDates, float64(i))
	}

	normalizedDates = Normalizer(normalizedDates)
	return normalizedDates
}

func NamesNormalizer(dates []string) []float64 {
	var normalizedDates []float64
	for i := 0; i <= len(dates); i++ {
		normalizedDates = append(normalizedDates, float64(i))
	}

	normalizedDates = Normalizer(normalizedDates)
	return normalizedDates
}

func Normalizer(in []float64) []float64 {
	out := make([]float64, 0, len(in))
	//min, max := MinMax(in)
	var min, max float64 = 0,1000
	switch min {
	case max:
		out = make([]float64, len(in), len(in))
	default:
		for v := range in {
			out = append(out, ((in[v] - min) / (max - min)))
		}
	}

	return out
}

func MinMax(in []float64) (float64, float64) {
	var max = in[0]
	var min = in[0]
	for _, value := range in {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Average(in []float64) float64 {
	total := 0.0
	for _, v := range in {
		total += v
	}
	res := total / float64(len(in))
	return res
}