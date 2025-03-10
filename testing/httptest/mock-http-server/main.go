package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := http.DefaultServeMux
	handler.HandleFunc("GET /data", getData())
	handler.HandleFunc("POST /data/{id}", newData())

	ts := httptest.NewServer(handler)
	defer ts.Close()
	client := ts.Client()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/data", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Is GET response correct", bytes.Equal(buf, []byte("data")))
	id := 123
	req, err = http.NewRequestWithContext(context.Background(), http.MethodPost, fmt.Sprintf("%s/data/%d", ts.URL, id), nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	buf, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Is POST response correct?", bytes.Equal(buf, []byte(fmt.Sprintf("id: %d created", id))))
}

func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "data")
	}
}

func newData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "id: %s created", id)
	}
}
