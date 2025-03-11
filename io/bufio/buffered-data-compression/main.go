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
	srcFile := "src.txt"
	dstFile := "dst.gz"
	f, err := os.Open(srcFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	srcData, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("failed to read from file: %w", err)
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to set offset: %w", err)
	}

	zf, err := os.Create(dstFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer zf.Close()
	gzipWriter := gzip.NewWriter(zf)

	_, err = io.Copy(gzipWriter, f)
	if err != nil {
		return fmt.Errorf("failed to copy from reader: %w", err)
	}
	gzipWriter.Close()
	f.Close()

	fd, err := os.Open(dstFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer fd.Close()
	gzipReader, err := gzip.NewReader(fd)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()
	dstData, err := io.ReadAll(gzipReader)
	if err != nil {
		return fmt.Errorf("failed to read from gzip reader: %w", err)
	}
	fmt.Println("src data == dst data?", bytes.Equal(srcData, dstData))
	fmt.Println(string(srcData))
	fmt.Println(string(dstData))
	return nil
}
