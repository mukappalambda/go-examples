package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Post struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

var (
	url = flag.String("url", "https://jsonplaceholder.typicode.com/posts", "url")
)

func main() {
	flag.Parse()
	if err := run(*url); err != nil {
		log.Fatal(err)
	}

}

func run(url string) error {
	// get request
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("making get request %q: %w", url, err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	var posts []Post

	if err := json.Unmarshal(b, &posts); err != nil {
		return fmt.Errorf("unmarshalling byte: %w", err)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Body)
	}

	// post request
	jsonStr := []byte(`{"userId": 1, "title": "post title", "body": "post body"}`)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return fmt.Errorf("making post request: %q; %w", url, err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		return fmt.Errorf("retrieving response: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}
	fmt.Println(string(body))
	return nil
}
