package main

import (
	"bytes"
	"fmt"
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
	res, err := client.Get(fmt.Sprintf("%s/redirect", ts.URL))
	_ = err
	defer res.Body.Close()
	buf, _ := io.ReadAll(res.Body)
	if !bytes.Equal(buf, []byte("latest resource")) {
		t.Fatalf("got: %s; want: %s\n", string(buf), "latest resource")
	}
	res, err = client.Get(fmt.Sprintf("%s/noredirect", ts.URL))
	_ = err
	defer res.Body.Close()
	buf, err = io.ReadAll(res.Body)
	_ = err
	if !bytes.Equal(buf, []byte("older resource")) {
		t.Fatalf("got: %s; want: %s\n", string(buf), "older resource")
	}
}
