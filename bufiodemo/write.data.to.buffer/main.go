package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	byt := []byte("Question: Do you love Go?\n")
	w := bufio.NewWriter(os.Stderr)
	_, err := w.Write(byt)
	if err != nil {
		log.Fatal(err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}

	s := "Answer: I love Go.\n"
	_, err = w.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
