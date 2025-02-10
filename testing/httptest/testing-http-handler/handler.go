package main

import (
	"fmt"
	"net/http"
)

func GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom-Header", "my-value")
		fmt.Fprint(w, "get all users")
	}
}
