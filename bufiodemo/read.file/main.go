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
	fmt.Println("Method 1: Read files using bufio.NewScanner()...")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Bytes())
	}
	fmt.Println("Method 2: Read files using bufio.NewReader()...")
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("error setting the offset: %s\n", err)
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("error reading from file: %s\n", err)
		}
		fmt.Print(line)
	}
}
