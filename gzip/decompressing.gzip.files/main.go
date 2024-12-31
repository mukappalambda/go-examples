package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("my-file.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	zr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
}
