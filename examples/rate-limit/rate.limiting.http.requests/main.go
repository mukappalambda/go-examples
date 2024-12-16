package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(1, 4)
	fmt.Println("rate is 1; burst is 4")
	handler := http.DefaultServeMux
	handler.HandleFunc("/data", handleData(limiter))
	ts := httptest.NewServer(handler)
	defer ts.Close()

	client := ts.Client()

	now := time.Now()
	var elapsed time.Duration
	for i := 0; i <= 20; i++ {
		elapsed = time.Since(now)
		res, err := client.Get(fmt.Sprintf("%s/data", ts.URL))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		buf, _ := io.ReadAll(res.Body)
		fmt.Printf("%02d-th request, %s, response: %s", i, elapsed.Truncate(time.Millisecond), string(buf))
		time.Sleep(200 * time.Millisecond)
	}
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
