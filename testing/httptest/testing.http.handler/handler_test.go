package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	ts := httptest.NewServer(GetUsers())
	defer ts.Close()
	client := ts.Client()
	res, _ := client.Get(ts.URL)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("want %d; got %d\n", http.StatusOK, res.StatusCode)
	}
	var want string
	want = "my-value"
	got := res.Header.Get("X-Custom-Header")
	if got != want {
		t.Fatalf("want: %q; got %q\n", want, got)
	}
	want = "get all users"
	buf, _ := io.ReadAll(res.Body)
	if !bytes.Equal(buf, []byte(want)) {
		t.Fatalf("want: %s; got: %s\n", want, string(buf))
	}
}
