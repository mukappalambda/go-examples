package main

import (
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
		w.Write([]byte("hi there\n"))
	}
}
