package main

import (
	"bytes"
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
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	var err error
	defer func() {
		err = zw.Close()
	}()
	if err != nil {
		return fmt.Errorf("error closing writer: %w", err)
	}
	data := []byte(`gzip is a file format and a software application used for file compression and decompression. The program was created by Jean-loup Gailly and Mark Adler as a free software replacement for the compress program used in early Unix systems, and intended for use by GNU (from which the "g" of gzip is derived). Version 0.1 was first publicly released on 31 October 1992, and version 1.0 followed in February 1993.`)
	if _, err := zw.Write(data); err != nil {
		return fmt.Errorf("error writing data to writer: %w", err)
	}
	zw.Close()
	zr, err := gzip.NewReader(&buf)
	if err != nil {
		return fmt.Errorf("error creating reader: %w", err)
	}
	defer zr.Close()
	if _, err := io.CopyN(os.Stdout, zr, 32768); err != nil {
		return fmt.Errorf("error copying to writer from reader: %w", err)
	}
	return nil
}
