package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

type MyRoundTripper struct {
	*log.Logger
}

func NewMyRoundTripper(logger *log.Logger) *MyRoundTripper {
	return &MyRoundTripper{
		Logger: logger,
	}
}

func (m *MyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	transport := http.DefaultTransport
	start := time.Now()
	res, err := transport.RoundTrip(r)
	elapsed := time.Since(start)
	m.Logger.Printf("method: %q; path: %q; elapsed time: %s\n", r.Method, r.URL.Path, elapsed)
	return res, err
}

var (
	_  http.RoundTripper = (*MyRoundTripper)(nil)
	ts *httptest.Server
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /fast", handleData(0))
	mux.HandleFunc("GET /slow", handleData(500*time.Millisecond))
	ts = httptest.NewServer(mux)
	defer ts.Close()
	log.Printf("test server is running on %q\n", ts.URL)
	ctx := context.Background()
	if err := fetchHTTPResponse(ctx, "/fast"); err != nil {
		ts.Close()
		return err
	}
	if err := fetchHTTPResponse(ctx, "/slow"); err != nil {
		ts.Close()
		return err
	}
	return nil
}

func fetchHTTPResponse(ctx context.Context, path string) error {
	client := DefaultClient()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL+path, nil)
	if err != nil {
		return fmt.Errorf("failed to create new request to %s: %s", path, err)
	}
	log.Printf("client is sending request: %s\n", req.URL)
	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		return fmt.Errorf("failed to send request: %s", err)
	}
	res.Body.Close()
	return nil
}

func DefaultClient() *http.Client {
	return NewClient(NewMyRoundTripper(log.Default()))
}

func NewClient(transport http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: transport,
	}
}

func handleData(delay time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		if _, err := fmt.Fprint(w, "data"); err != nil {
			http.Error(w, "error responding", http.StatusInternalServerError)
		}
	}
}
