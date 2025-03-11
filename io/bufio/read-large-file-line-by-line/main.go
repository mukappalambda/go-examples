package main

import (
	"bufio"
	"bytes"
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
	f, err := os.CreateTemp("", "temp_large_file.txt")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	count := 10000
	if _, err := f.Write(bytes.Repeat([]byte("hello?\n"), count)); err != nil {
		f.Close()
		return fmt.Errorf("error writing to temp file: %w", err)
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("error setting the offset: %w", err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_ = scanner.Text()
	}
	fmt.Println("read the large file successfully.")
	return nil
}
