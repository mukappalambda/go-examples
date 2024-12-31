package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	defer func() {
		if err := zw.Close(); err != nil {
			log.Fatalf("error closing writer: %s\n", err)
		}
	}()
	data := []byte(`gzip is a file format and a software application used for file compression and decompression. The program was created by Jean-loup Gailly and Mark Adler as a free software replacement for the compress program used in early Unix systems, and intended for use by GNU (from which the "g" of gzip is derived). Version 0.1 was first publicly released on 31 October 1992, and version 1.0 followed in February 1993.`)
	if _, err := zw.Write(data); err != nil {
		log.Fatalf("error writing data to writer: %s\n", err)
	}
	zw.Close()
	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatalf("error creating reader: %s\n", err)
	}
	defer zr.Close()
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatalf("error copying to writer from reader: %s\n", err)
	}
}
