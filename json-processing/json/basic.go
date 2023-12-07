package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	bytBool, err := json.Marshal(true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytBool))

	var datBool bool
	if err := json.Unmarshal(bytBool, &datBool); err != nil {
		log.Fatal(err)
	}
	fmt.Println(datBool)

	bytInt, err := json.Marshal(123)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytInt))

	var datInt int
	if err := json.Unmarshal(bytInt, &datInt); err != nil {
		log.Fatal(err)
	}
	fmt.Println(datInt)

	bytStr, err := json.Marshal("foobar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytStr))

	var datStr string
	if err := json.Unmarshal(bytStr, &datStr); err != nil {
		log.Fatal(err)
	}
	fmt.Println(datStr)

	sliceStr := []string{"alex", "bob", "mark"}
	bytSliceStr, err := json.Marshal(sliceStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytSliceStr))

	var datSliceStr []string
	if err := json.Unmarshal(bytSliceStr, &datSliceStr); err != nil {
		log.Fatal(err)
	}
	fmt.Println(datSliceStr)

	mapStrFloat := map[string]float32{"alex": 1.11, "bob": 2.22, "mark": 3.33}
	bytMap, err := json.Marshal(mapStrFloat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytMap))

	var datMap map[string]float32
	if err := json.Unmarshal(bytMap, &datMap); err != nil {
		log.Fatal(err)
	}
	fmt.Println(datMap)
}
