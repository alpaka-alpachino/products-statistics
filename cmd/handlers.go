package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type IndexStats struct {
	HotRises           []string
	RegionsForProducts map[string]string
	Selected           string
	Prediction         string
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("front/index.html", "front/header.html", "front/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	stats := IndexStats{HotRises: hotRises,
		RegionsForProducts: regionsForSelected,
		Selected:           selected,
		Prediction:         predictionForSelected,
	}

	err = t.ExecuteTemplate(w, "index", stats)
	if err != nil {
		return
	}
}

func handleFunc() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
