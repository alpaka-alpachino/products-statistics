package statistics

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/web-site/storage"
	"os"
	"time"
)

const sourceFormat = "host=localhost port=5432 user=%v password=%v dbname=%v sslmode=disable"

func GetHotRises()[]string{
	var hotRises []string
	var others []string
	var interest = storage.All
	for v := range interest {
		_, _, firstPrice, _, lastPrice, _, _, avg:= GetStatForProduct(interest[v])
		status,change := SimpleTrendCheck(firstPrice,avg,lastPrice)
		if status == "Hot Rise"{
			hotRises = append(hotRises, fmt.Sprintf(`%s:  %v`,interest[v],change))
		}else {
			others =  append(others, fmt.Sprintf(`%s:  %v`,interest[v],change))
		}
	}
	return hotRises
}

func GetStatForProduct(prod string) (string, time.Time, float64, time.Time, float64, float64, float64, float64) {
	//connect to db
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

	//retrieve data from db
	res, err := database.Query("SELECT price, date FROM prices WHERE product = $1 ORDER BY date ASC", prod)
	if err != nil {
		panic(err)
	}

	var prices []float64
	var dates []time.Time

	for res.Next() {
		var model = new(storage.Model)
		err = res.Scan(&model.Price, &model.Date)
		if err != nil {
			panic(err)
		}

		storage.Statistics = append(storage.Statistics, *model)
		prices = append(prices, model.Price)
		dates = append(dates, model.Date)
	}

	prices = prices[:len(prices)-1] //todo fix in scraper
	min, max := MinMax(prices)
	avg := Average(prices)

	if len(dates) > 0 && len(dates) < 50 {
		dates = dates[:len(dates)-1]
		if len(dates) > 0 {
			return prod, dates[0], prices[0], dates[len(dates)-1], prices[len(prices)-1], min, max, avg
		}
	}

	return `unclear`, time.Time{}, 0.0, time.Time{}, 0.0, 0.0, 0.0, 0.0
}

func SimpleTrendCheck(initial, average, final float64) (string, float64) {
	var wholeChange = final - initial
	var hotChanges = final - average
	if wholeChange > 0 {
		if hotChanges > wholeChange {
			return "Hot Rise", wholeChange
		} else {
			return "Rise", wholeChange
		}
	} else if wholeChange < 0 {
		if hotChanges < wholeChange {
			return "Hot Cheapening", wholeChange
		} else {
			return "Cheapening", wholeChange
		}
	} else {
		return "Stable", 0
	}
}
