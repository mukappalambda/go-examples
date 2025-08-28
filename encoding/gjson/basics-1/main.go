package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tidwall/gjson"
)

type Data map[string]interface{}

func main() {
	byt, err := os.ReadFile("./example.json")
	if err != nil {
		log.Fatal(err)
	}
	cnt := gjson.GetBytes(byt, "#")
	secondLatitude := gjson.GetBytes(byt, "1.latitude")
	firstFriendNames := gjson.GetBytes(byt, "0.friends.#.name")
	filteredNames := gjson.GetBytes(byt, "#(age<30)#.name")
	fmt.Println(cnt.Int())
	fmt.Println("Second latitude:", secondLatitude.Float())
	for _, name := range firstFriendNames.Array() {
		fmt.Println("name:", name.String())
	}
	fmt.Printf("%+v", filteredNames)
}
