package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type myRoundTripper struct{}

func (m *myRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	transport := http.DefaultTransport
	start := time.Now()
	res, err := transport.RoundTrip(r)
	elapsed := time.Since(start)
	log.Printf("method: %q; path: %q; elapsed time: %s\n", r.Method, r.URL.Path, elapsed)
	return res, err
}

var _ http.RoundTripper = (*myRoundTripper)(nil)

func main() {
	http.HandleFunc("GET /fast", handleData(0))
	http.HandleFunc("GET /slow", handleData(500*time.Millisecond))
	ts := httptest.NewServer(nil)
	defer ts.Close()
	client := NewClient()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/fast", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/slow", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()
}

func NewClient() *http.Client {
	client := &http.Client{
		Transport: &myRoundTripper{},
	}
	return client
}

func handleData(delay time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		if _, err := fmt.Fprint(w, "data"); err != nil {
			http.Error(w, "error responding", http.StatusInternalServerError)
		}
	}
}
