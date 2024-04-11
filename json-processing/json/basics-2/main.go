package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Post struct {
	Userid int    `json:"userId,omitempty"`
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	b, err := os.ReadFile("./posts.json")
	if err != nil {
		return err
	}

	var posts []Post

	if err := json.Unmarshal(b, &posts); err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Userid, post.Id, post.Title)
	}
	return nil
}
