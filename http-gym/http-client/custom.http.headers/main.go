package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

var apiKey = "my-key"

func main() {
	http.HandleFunc("GET /data", HeaderValidator(handleData()))
	ts := httptest.NewServer(nil)
	defer ts.Close()

	client := &http.Client{
		Timeout: time.Second,
	}
	resourcePath := fmt.Sprintf("%s/data", ts.URL)

	b, err := requestData(client, resourcePath)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("client does not set the custom header, so response is %q\n", string(b))

	b, err = requestData(client, resourcePath, "wrong-header", "wrong-value")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("client set the wrong header, so response is %q\n", string(b))

	b, err = requestData(client, resourcePath, "X-API-KEY", apiKey)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("client set the correct header, so response is %q\n", string(b))
}

func requestData(client *http.Client, resourcePath string, args ...string) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, resourcePath, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new request: %s", err)
	}
	if len(args) == 2 {
		key := args[0]
		value := args[1]
		req.Header.Set(key, value)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP GET request: %s", err)
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}
	return b, nil
}

func HeaderValidator(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-KEY")
		if key == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			log.Printf("[client@%s] \"bad request\" %d\n", r.RemoteAddr, http.StatusBadRequest)
			return
		}
		if key != apiKey {
			http.Error(w, "invalid api key", http.StatusBadRequest)
			log.Printf("[client@%s] \"invalid api key\" %d\n", r.RemoteAddr, http.StatusBadRequest)
		}
		log.Printf("[client@%s] validated\n", r.RemoteAddr)
		h(w, r)
	}
}

func handleData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, "data"); err != nil {
			http.Error(w, "error responding", http.StatusInternalServerError)
		}
	}
}
