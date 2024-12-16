package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

// This server echoes any message it receives from the client.
func main() {
	addr := ":8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("failed to listen at %s, %s\n", addr, err)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept the connection: %s\n", err)
			os.Exit(1)
		}
		go func(conn net.Conn) {
			defer conn.Close()
			reader := bufio.NewReader(conn)
			_, err = io.Copy(conn, reader)
			if err != nil {
				log.Printf("failed to copy: %s\n", err)
				os.Exit(1)
			}
		}(conn)
	}
}
