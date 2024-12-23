package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var i int
	for i = 0; s.Scan(); i++ {
		fmt.Printf("%s\n", s.Bytes())
	}
}
