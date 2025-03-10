package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func handleData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "data")
	}
}

func handleGreet(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, msg)
	}
}

func handleSlow(duration time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		fmt.Fprint(w, "ok")
	}
}
