package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFooMiddleware(t *testing.T) {
	handler := http.DefaultServeMux
	handler.Handle("GET /middleware", FooMiddleware(MyHandler()))
	ts := httptest.NewServer(handler)
	defer ts.Close()
	client := ts.Client()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/middleware", nil)
	if err != nil {
		t.Fail()
	}
	res, err := client.Do(req)
	_ = err
	defer res.Body.Close()
	got := res.Header.Get("X-Custom-Middleware")
	want := "foo middleware"
	if !bytes.Equal([]byte(got), []byte(want)) {
		t.Fatalf("got: %s; want: %s\n", got, want)
	}
}
