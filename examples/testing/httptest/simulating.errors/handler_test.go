package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler(t *testing.T) {
	ts := httptest.NewServer(MyHandler())
	defer ts.Close()
	client := ts.Client()
	res, err := client.Get(ts.URL)
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
