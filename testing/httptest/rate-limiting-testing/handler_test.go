package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimitedHandler(t *testing.T) {
	r := 190
	b := 10
	ts := httptest.NewServer(RateLimitedHandler(r, b))
	defer ts.Close()
	client := ts.Client()

	num := r + b
	for i := 0; i < num; i++ {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
		if err != nil {
			t.Fail()
		}
		res, err := client.Do(req)
		_ = err
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Fatalf("got %d\n", res.StatusCode)
		}
		time.Sleep(5 * time.Millisecond)
	}
	time.Sleep(20 * time.Millisecond)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fail()
	}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("got %d\n", res.StatusCode)
	}
}
