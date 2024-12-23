package main

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimitedHandler(r, b int) http.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(r), b)
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("hi there\n")); err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			log.Printf("Error writing response: %s\n", err)
		}
	}
}
