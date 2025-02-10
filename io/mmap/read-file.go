package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/mmap"
)

var fileName = flag.String("file", "", "path to file")

func main() {
	flag.Parse()
	if len(*fileName) == 0 {
		fmt.Print("forgot to pass filename?")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fileContent, err := fetchFileContent(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(fileContent))
}

func fetchFileContent(fileName string) ([]byte, error) {
	r, err := mmap.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file using mmap: %s", err.Error())
	}
	got := make([]byte, r.Len())
	_, err = r.ReadAt(got, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %s", err.Error())
	}
	return got, nil
}
