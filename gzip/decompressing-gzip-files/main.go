package main

import (
	"compress/gzip"
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
	f, err := os.Open("my-file.txt.gz")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()
	zr, err := gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("failed to create reader: %w", err)
	}
	defer zr.Close()
	if _, err := io.CopyN(os.Stdout, zr, 4096); err != nil {
		return fmt.Errorf("failed to copy from gzip reader: %w", err)
	}
	return nil
}
