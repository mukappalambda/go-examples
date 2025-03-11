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
	w := bufio.NewWriter(os.Stdout)
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read: %w", err)
		}
		_, err = w.Write(b)
		if err != nil {
			return fmt.Errorf("failed to write bytes: %w", err)
		}
		err = w.Flush()
		if err != nil {
			return fmt.Errorf("failed to flush: %w", err)
		}
	}
	return nil
}
