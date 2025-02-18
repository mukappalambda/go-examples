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

var (
	_  http.RoundTripper = (*myRoundTripper)(nil)
	ts *httptest.Server
)

func main() {
	http.HandleFunc("GET /fast", handleData(0))
	http.HandleFunc("GET /slow", handleData(500*time.Millisecond))
	ts = httptest.NewServer(nil)
	if err := run(); err != nil {
		ts.Close()
		log.Fatal(err)
	}
	ts.Close()
}

func run() error {
	ctx := context.Background()
	if err := fetchHTTPResponse(ctx, "/fast"); err != nil {
		return err
	}
	if err := fetchHTTPResponse(ctx, "/slow"); err != nil {
		return err
	}
	return nil
}

func fetchHTTPResponse(ctx context.Context, path string) error {
	client := NewClient()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL+path, nil)
	if err != nil {
		return fmt.Errorf("failed to create new request to %s: %s", path, err)
	}
	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		return fmt.Errorf("failed to send request: %s", err)
	}
	res.Body.Close()
	return nil
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
