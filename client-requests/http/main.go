package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	// get request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the response body.")
	}

	var posts []Post

	if err := json.Unmarshal(b, &posts); err != nil {
		fmt.Println("Failed to unmarshal the request body.")
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Body)
	}

	// post request
	jsonStr := []byte(`{"userId": 1, "title": "post title", "body": "post body"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
