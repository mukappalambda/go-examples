package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	num := 0
	for {
		_, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		num++
	}
	fmt.Printf("Number of lines in main.go: %d\n", num)
}
