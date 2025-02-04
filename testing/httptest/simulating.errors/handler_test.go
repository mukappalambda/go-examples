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
	ts := httptest.NewServer(MyHandler())
	defer ts.Close()
	client := ts.Client()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fail()
	}
	res, err := client.Do(req)
	_ = err
	defer res.Body.Close()
	got := res.StatusCode
	want := http.StatusInternalServerError
	if got != want {
		t.Fatalf("got: %d; want: %d\n", got, want)
	}
	buf, err := io.ReadAll(res.Body)
	_ = err
	if !bytes.Equal(buf, []byte("server experienced error\n")) {
		t.Fatal("unexpected server response")
	}
}
