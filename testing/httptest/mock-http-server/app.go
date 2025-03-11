package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

func run() error {
	ts := NewHTTPServer()
	defer ts.Close()
	client := ts.Client()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/data", nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP GET request: %w", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Println("Is GET response correct", bytes.Equal(buf, []byte("data")))
	id := 123
	req, err = http.NewRequestWithContext(context.Background(), http.MethodPost, fmt.Sprintf("%s/data/%d", ts.URL, id), nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP POST request: %w", err)
	}
	res, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer res.Body.Close()
	buf, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Println("Is POST response correct?", bytes.Equal(buf, []byte(fmt.Sprintf("id: %d created", id))))
	return nil
}
