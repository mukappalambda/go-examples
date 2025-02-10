package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	srcFile := "src.txt"
	dstFile := "dst.gz"
	f, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	srcData, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	zf, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer zf.Close()
	gzipWriter := gzip.NewWriter(zf)

	_, err = io.Copy(gzipWriter, f)
	if err != nil {
		log.Fatal(err)
	}
	gzipWriter.Close()
	f.Close()

	fd, err := os.Open(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	gzipReader, err := gzip.NewReader(fd)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()
	dstData, err := io.ReadAll(gzipReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("src data == dst data?", bytes.Equal(srcData, dstData))
	fmt.Println(string(srcData))
	fmt.Println(string(dstData))
}
