package main

import (
	"fmt"
	"net/http"
)

func FooMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom-Middleware", "foo middleware")
		h.ServeHTTP(w, r)
	})
}

func MyHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "my handler")
		})
}
