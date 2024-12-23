package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("large_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	cnt := 1
	for scanner.Scan() {
		if cnt > 3 {
			break
		}
		fmt.Printf("Read contents: %s\n", scanner.Text())
		cnt = cnt + 1
	}
}
