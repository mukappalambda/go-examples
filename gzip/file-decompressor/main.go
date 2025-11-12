package main

import (
	"compress/gzip"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	ErrNoFile = errors.New("missing input file")

	input = flag.String("in", "", "input file")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(*input) == 0 {
		flag.PrintDefaults()
		return ErrNoFile
	}
	f, err := os.Open(*input)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()
	zr, err := gzip.NewReader(f)
	if err != nil {
		return fmt.Errorf("failed to create reader: %w", err)
	}
	defer func() {
		if err := zr.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to close gzip reader: %s", err)
		}
	}()
	if _, err := io.CopyN(os.Stdout, zr, 4096); err != nil && err != io.EOF {
		return fmt.Errorf("failed to copy from gzip reader: %w", err)
	}
	return nil
}
