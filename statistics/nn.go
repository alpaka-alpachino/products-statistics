package statistics

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/web-site/storage"
)

const (
	selectQuery         = "SELECT date, price FROM prices WHERE product = $1 ORDER BY date ASC"
	min, max    float64 = 0, 1000
)

type NNStatsModel struct {
	Rows []Rows
}

type Rows struct {
	Dates  time.Time
	Prices float64
}

// GetTrainingData returns training data from given targets and stats
func GetTrainingData(targets, stats []NNStatsModel, rawsLen int) [][][]float64 {
	var targetPrices []float64

	for i := 0; i < rawsLen; i++ {
		for _, v := range targets {
			targetPrices = append(targetPrices, v.Rows[i].Prices)
		}
	}

	targetPrices = Normalizer(targetPrices)

	for _, stat := range stats {
		prices := make([]float64, 0, len(stat.Rows))

		for _, raw := range stat.Rows {
			prices = append(prices, raw.Prices)
		}

		normalizedPrices := Normalizer(prices)

		for i, v := range normalizedPrices {
			stat.Rows[i].Prices = v
		}
	}

	var normalizedSideProds [][]float64

	for i := 0; i < rawsLen; i++ {
		var normalizedSidePrices []float64

		for _, v := range stats {
			normalizedSidePrices = append(normalizedSidePrices, v.Rows[i].Prices)
		}

		normalizedSideProds = append(normalizedSideProds, normalizedSidePrices)

	}

	var trainingData [][][]float64

	targetPrices = targetPrices[1:]
	normalizedSideProds = normalizedSideProds[:len(normalizedSideProds)-1]

	for i, v := range normalizedSideProds {
		var sideProdsColumns [][]float64

		sideProdsColumns = append(sideProdsColumns, v)
		sideProdsColumns = append(sideProdsColumns, []float64{targetPrices[i]})

		trainingData = append(trainingData, sideProdsColumns)
	}

	return trainingData
}

// GetStatsForNN retrieves data and len by keys,
// should accept at least 2 product types (keys) as input
func GetStatsForNN(prods []string) ([]NNStatsModel, int, error) {
	var rowsLen int
	s := fmt.Sprintf(sourceFormat, os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DB"))
	database, err := sql.Open("postgres", s)
	if err != nil {
		return nil, 0, err
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Println(err)
		}
	}(database)

	var prodValues []interface{}
	for v := range prods {
		prodValues = append(prodValues, prods[v])
	}

	set := make([]NNStatsModel, 0, 0)
	for p := range prodValues {
		res, err := database.Query(selectQuery, prodValues[p])
		if err != nil {
			return nil, 0, err
		}

		var raws []Rows
		for res.Next() {
			var model = new(storage.ProductModel)
			err = res.Scan(&model.Date, &model.Price)
			if err != nil {
				return nil, 0, err
			}

			storage.Statistics = append(storage.Statistics, *model)
			raws = append(raws, Rows{
				Dates:  model.Date,
				Prices: model.Price,
			})

			rowsLen = len(raws) - 1 // may be changed in scraper
		}
		set = append(set, NNStatsModel{Rows: raws[:len(raws)-1]})
	}

	return set, rowsLen, nil
}

// Normalizer normalizes data
func Normalizer(in []float64) []float64 {
	out := make([]float64, 0, len(in))
	switch min {
	case max:
		out = make([]float64, len(in), len(in))
	default:
		for v := range in {
			out = append(out, (in[v]-min)/(max-min))
		}
	}

	return out
}
