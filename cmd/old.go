package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(`front/index.html`, `front/design/header.html`, `front/design/footer.html`)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	err = tmp.ExecuteTemplate(w, `index`, nil)
	if err != nil {
		return
	}
}

func handleFunc() {
	http.HandleFunc(`/`, index)
	err := http.ListenAndServe(`:8000`, nil)
	if err != nil {
		return
	}
}
