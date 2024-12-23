package main

import (
	"fmt"
	"net/http"
)

func HandleOlderResource(redirect bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if redirect {
			http.Redirect(w, r, "/latest", http.StatusPermanentRedirect)
			return
		}
		fmt.Fprint(w, "older resource")
	}
}

func HandleLatestResource() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "latest resource")
	}
}
