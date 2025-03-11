package main

import (
	"fmt"
	"net/http"
)

func defaultMux() *http.ServeMux {
	mux := http.DefaultServeMux
	mux.HandleFunc("GET /", handleRoot())
	return mux
}

func handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "root")
	}
}
