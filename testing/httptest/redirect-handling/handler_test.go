package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler(t *testing.T) {
	handler := http.DefaultServeMux
	handler.HandleFunc("/redirect", HandleOlderResource(true))
	handler.HandleFunc("/noredirect", HandleOlderResource(false))
	handler.HandleFunc("/latest", HandleLatestResource())
	ts := httptest.NewServer(handler)
	defer ts.Close()
	client := ts.Client()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/redirect", nil)
	if err != nil {
		t.Fail()
	}
	res, err := client.Do(req)
	_ = err
	defer res.Body.Close()
	buf, _ := io.ReadAll(res.Body)
	if !bytes.Equal(buf, []byte("latest resource")) {
		t.Fatalf("got: %s; want: %s\n", string(buf), "latest resource")
	}
	req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/noredirect", nil)
	if err != nil {
		t.Fail()
	}
	res, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	_ = err
	defer res.Body.Close()
	buf, err = io.ReadAll(res.Body)
	_ = err
	if !bytes.Equal(buf, []byte("older resource")) {
		t.Fatalf("got: %s; want: %s\n", string(buf), "older resource")
	}
}
