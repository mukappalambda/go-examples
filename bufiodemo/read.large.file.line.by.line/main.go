package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "temp_large_file.txt")
	if err != nil {
		log.Fatalf("error creating temp file: %s\n", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	count := 10000
	if _, err := f.Write(bytes.Repeat([]byte("hello?\n"), count)); err != nil {
		f.Close()
		log.Fatalf("error writing to temp file: %s\n", err)
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("error setting the offset: %s\n", err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_ = scanner.Text()
	}
	fmt.Println("read the large file successfully.")
}
