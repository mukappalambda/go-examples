package main

import (
	"bufio"
	"bytes"
	"fmt"
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

	buf := bytes.NewBuffer([]byte{})
	writer := bufio.NewWriter(buf)
	_, _ = writer.Write([]byte("Go is great.\n"))
	writer.Flush()
	fmt.Print(buf.String())
}
