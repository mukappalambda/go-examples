package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Post struct {
  Userid int `json:"userId"`
  Id int `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
}

func main()  {
  b, err := ioutil.ReadFile("./posts.json")

  if err != nil {
    log.Fatal(err)
  }

  var posts []Post

  if err := json.Unmarshal(b, &posts); err != nil {
    fmt.Println(err)
  }

  for _, post := range posts {
    fmt.Println(post.Userid, post.Id, post.Title)
  }
}