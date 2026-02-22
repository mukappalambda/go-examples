package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	jsonData := []byte(`{
  "userId": 1,
  "id": 1,
  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
}`)
	var data map[string]any
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
