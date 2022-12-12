package main

import "fmt"

func main() {
	hashmap := make(map[int]struct{})

	str := []int{1, 2, 3, 4, 5}

	for _, item := range str {
		hashmap[item] = struct{}{}
	}

	if _, ok := hashmap[1]; ok {
		fmt.Println("1 is in hashmap")
	}

}
