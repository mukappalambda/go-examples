package main

import (
	"fmt"
	"net/http"
)

func defaultHandler() http.Handler {
	handler := http.DefaultServeMux
	handler.HandleFunc("GET /data", getData())
	handler.HandleFunc("POST /data/{id}", newData())
	return handler
}

func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "data")
	}
}

func newData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "id: %s created", id)
	}
}
