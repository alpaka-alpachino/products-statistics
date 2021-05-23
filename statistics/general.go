package statistics

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/web-site/storage"
	"log"
	"os"
	"strings"
	"time"
)

const (
	sourceFormat = "host=localhost port=5432 user=%v password=%v dbname=%v sslmode=disable"

	HotRises      = "Hot Rise"
	Rises         = "Rise"
	HotCheapening = "Hot Cheapening"
	Cheapening    = "Cheapening"
	Stable        = "Stable"
)

var errEmptyDate = errors.New("empty date slice")

type ProductStats struct {
	name         string
	firstDate    time.Time
	firstPrice   float64
	lastDate     time.Time
	lastPrice    float64
	minPrice     float64
	maxPrice     float64
	averagePrice float64
}

// GetChanges gets prices and its change if it changes during period
func GetChanges(inputs []string, requestedStatus string) ([]string, error) {
	var hotRises, others []string
	for v := range inputs {
		productStats, err := GetStatForProduct(inputs[v])
		if err != nil {
			return nil, err
		}
		status, change := SimpleTrendCheck(productStats.firstPrice, productStats.averagePrice, productStats.lastPrice)
		if status == requestedStatus {
			hotRises = append(hotRises, fmt.Sprintf("%s -->  %.2f;", inputs[v], change))
		} else {
			others = append(others, fmt.Sprintf("%s -->  %.2f;", inputs[v], change))
		}
	}
	return hotRises, nil
}

// GetStatForProduct retrieves data for requested product from db
func GetStatForProduct(prod string) (ProductStats, error) {
	s := fmt.Sprintf(sourceFormat, os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DB"))
	database, err := sql.Open("postgres", s)
	if err != nil {
		return ProductStats{}, err
	}
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Println(err)
		}
	}(database)

	res, err := database.Query(selectQuery, prod)
	if err != nil {
		return ProductStats{}, err
	}

	var prices []float64
	var dates []time.Time

	for res.Next() {
		var model = new(storage.ProductModel)
		err = res.Scan(&model.Date, &model.Price)
		if err != nil {
			return ProductStats{}, err
		}

		storage.Statistics = append(storage.Statistics, *model)
		prices = append(prices, model.Price)
		dates = append(dates, model.Date)
	}

	prices = prices[:len(prices)-1] //can be changed in scraper
	min, max := MinMax(prices)
	avg := Average(prices)

	if len(dates) > 0 {
		if len(dates) > 50 { // use this check if dont need different products
			return ProductStats{}, nil
		}
		dates = dates[:len(dates)-1]
		if len(dates) > 0 {
			return ProductStats{
				name:         prod,
				firstDate:    dates[0],
				firstPrice:   prices[0],
				lastDate:     dates[len(dates)-1],
				lastPrice:    prices[len(prices)-1],
				minPrice:     min,
				maxPrice:     max,
				averagePrice: avg,
			}, nil
		}
	}

	return ProductStats{}, errEmptyDate
}

// SimpleTrendCheck checks if there are changes in price during period
func SimpleTrendCheck(initial, average, final float64) (string, float64) {
	var wholeChange = final - initial
	var hotChanges = final - average
	if wholeChange > 0 {
		if hotChanges > wholeChange {
			return HotRises, wholeChange
		} else {
			return Rises, wholeChange
		}
	} else if wholeChange < 0 {
		if hotChanges < wholeChange {
			return HotCheapening, wholeChange
		} else {
			return Cheapening, wholeChange
		}
	} else {
		return Stable, 0
	}
}

// GetRegionsForProducts returns list of regions with its land value related to input products
func GetRegionsForProducts(inputs []string) map[string]string {
	var regionsForProducts = make(map[string]string)
	for p := range inputs {
		for k, v := range storage.ProductRegions {
			if strings.Contains(inputs[p], k) {
				for vv := range v {
					regionsForProducts[inputs[p]] = fmt.Sprintf("%s %s (%d)", regionsForProducts[inputs[p]], v[vv].Name, v[vv].Index)
				}
			}
		}
	}
	return regionsForProducts
}

// MinMax gets minimal and maximum values from array
func MinMax(inputs []float64) (float64, float64) {
	var max = inputs[0]
	var min = inputs[0]
	for _, value := range inputs {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	return min, max
}

// Average gets average value from array
func Average(inputs []float64) float64 {
	total := 0.0
	for _, v := range inputs {
		total += v
	}
	res := total / float64(len(inputs))
	return res
}
