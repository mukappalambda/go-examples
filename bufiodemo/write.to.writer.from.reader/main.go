package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		_, err = w.Write(b)
		if err != nil {
			log.Fatal(err)
		}
		err = w.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}
}
