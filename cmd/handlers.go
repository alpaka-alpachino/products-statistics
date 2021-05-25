package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
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
	t, err := template.ParseFiles("front/index.html", "front/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	stats := IndexStats{
		HotRises:           hotRises,
		RegionsForProducts: regionsForSelected,
		Selected:           selected,
		Prediction:         predictionForSelected,
	}

	err = t.ExecuteTemplate(w, "index", stats)
	if err != nil {
		return
	}
}

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	addr := *flag.String("address", ":8080", "address of server")
	flag.Parse()

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")

	return &Server{
		&http.Server{
			Addr:    addr,
			Handler: limit(rtr),
		},
	}
}

func (ws *Server) Run() error {
	return ws.server.ListenAndServe()
}
