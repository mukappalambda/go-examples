package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open("./main.go")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
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
			return fmt.Errorf("error reading string from file: %w", err)
		}
		num++
	}
	fmt.Printf("Number of lines in main.go: %d\n", num)
	return nil
}
