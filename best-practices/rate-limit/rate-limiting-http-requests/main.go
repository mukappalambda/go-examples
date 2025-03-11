package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	limiter := rate.NewLimiter(1, 4)
	fmt.Println("rate is 1; burst is 4")
	handler := http.DefaultServeMux
	handler.HandleFunc("/data", handleData(limiter))
	ts := httptest.NewServer(handler)
	defer ts.Close()

	client := ts.Client()

	now := time.Now()
	var elapsed time.Duration
	ctx := context.Background()
	for i := 0; i <= 20; i++ {
		elapsed = time.Since(now)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/data", ts.URL), nil)
		if err != nil {
			return err
		}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		buf, _ := io.ReadAll(res.Body)
		fmt.Printf("%02d-th request, %s, response: %s", i, elapsed.Truncate(time.Millisecond), string(buf))
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}

func handleData(limiter *rate.Limiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "rate limiting", http.StatusTooManyRequests)
			return
		}
		fmt.Fprintln(w, "data")
	}
}
