package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("error creating a connection: %s\n", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Printf("error reading input: %s\n", err)
			break
		}
		if text == "\n" {
			continue
		}
		if _, err := conn.Write([]byte(text)); err != nil {
			log.Fatalf("error writing to connection: %s\n", err)
		}
		serverResponse, err := connReader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(serverResponse)
	}
}
