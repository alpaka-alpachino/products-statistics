package storage

import "time"

type ProductModel struct {
	Date    time.Time `json:"date"`
	Product string    `json:"product"`
	Price   float64   `json:"price"`
}

var Statistics []ProductModel
