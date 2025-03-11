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
	fmt.Println("Method 1: Read files using bufio.NewScanner()...")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Bytes())
	}
	fmt.Println("Method 2: Read files using bufio.NewReader()...")
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("error setting the offset: %w", err)
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading from file: %w", err)
		}
		fmt.Print(line)
	}
	return nil
}
