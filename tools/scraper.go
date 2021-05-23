package tools

import (
	"database/sql"
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL      = `https://index.minfin.com.ua/ua/`
	sourceFormat = "host=localhost port=5432 user=%v password=%v dbname=%v sslmode=disable"
	dataLayout   = "02.01.2006"
	sqlStatement = "INSERT INTO prices(date, product, price) VALUES($1,$2,$3)"
)

// ScrapeDaily scrapes daily info from https://index.minfin.com.ua/ua/ and put it into db
func ScrapeDaily() {
	collectGroups := colly.NewCollector()
	collectItem := colly.NewCollector()
	collectInfo := colly.NewCollector()

	s := fmt.Sprintf(sourceFormat, os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DB"))

	database, err := sql.Open("postgres", s)
	if err != nil {
		panic(err)
	}

	defer database.Close()

	collectGroups.OnHTML("#idx-content > ul > li > a", func(e *colly.HTMLElement) {
		link := baseURL + e.Attr("href")
		collectItem.Visit(link)
	})

	collectItem.OnHTML("#idx-content > div > table > tbody > tr > td > a", func(e *colly.HTMLElement) {
		link := baseURL + e.Attr("href")
		collectInfo.Visit(link)
	})

	var product string
	var date time.Time
	var price float64

	collectInfo.OnHTML(`#idx-content > h3`, func(e *colly.HTMLElement) {
		data := e.Text
		data = strings.ReplaceAll(data, `Бакалія: `, ``)
		product = data
	})

	collectInfo.OnHTML(`#idx-content > div > table > tbody > tr:nth-child(2) > th > big`, func(e *colly.HTMLElement) {
		data := strings.ReplaceAll(e.Text, `,`, `.`)
		price, err = strconv.ParseFloat(data, 64)
		if err != nil {
			panic(err)
		}

		date = time.Now()
		fmt.Println(date, product, price)
		insert, err := database.Query(sqlStatement, date, product, price)
		if err != nil {
			panic(err)
		}
		defer insert.Close()
	})

	collectGroups.Visit(`https://index.minfin.com.ua/ua/markets/wares/prods/`)
}

// ScrapeMonthly scrapes monthly info from https://index.minfin.com.ua/ua/ and put it into db
func ScrapeMonthly() {
	collectGroups := colly.NewCollector()
	collectItem := colly.NewCollector()
	collectInfo := colly.NewCollector()

	s := fmt.Sprintf(sourceFormat, os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DB"))

	database, err := sql.Open("postgres", s)
	if err != nil {
		panic(err)
	}

	defer database.Close()

	collectGroups.OnHTML("#idx-content > ul > li > a", func(e *colly.HTMLElement) {
		link := baseURL + e.Attr("href")
		collectItem.Visit(link)
	})

	collectItem.OnHTML("#idx-content > div > table > tbody > tr > td > a", func(e *colly.HTMLElement) {
		link := baseURL + e.Attr("href")
		collectInfo.Visit(link)
	})

	var product string
	var date time.Time
	var price float64

	collectInfo.OnHTML(`#idx-content > h3`, func(e *colly.HTMLElement) {
		data := e.Text
		data = strings.ReplaceAll(data, `Бакалія: `, ``)
		product = data
	})

	collectInfo.OnHTML(`#idx-content > div.prodsdataview > table > tbody > tr`, func(e *colly.HTMLElement) {
		e.ForEach(`td:nth-child(1)`, func(_ int, inner *colly.HTMLElement) {
			data := inner.Text
			date, err = time.Parse(dataLayout,data)
			if err != nil {
				panic(err)
			}
		})

		e.ForEach(` td:nth-child(3)`, func(_ int, inner *colly.HTMLElement) {
			data := strings.ReplaceAll(inner.Text,`,`,`.`)
			price, err = strconv.ParseFloat(data,64)
			if err != nil {
				panic(err)
			}
		})

		insert, err := database.Query(sqlStatement,date,product,price)
		if err != nil {
			panic(err)
		}
		defer insert.Close()
	})

	collectGroups.Visit(`https://index.minfin.com.ua/ua/markets/wares/prods/`)
}


