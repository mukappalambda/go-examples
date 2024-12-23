package main

import "net/http"

func MyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "server experienced error", http.StatusInternalServerError)
	}
}
