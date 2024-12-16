package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	port := 8080
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server is listening at %s\n", ln.Addr().String())
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to accept tcp connection: %s\n", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fileName := fmt.Sprintf("%s/tmp_file.txt", dir)
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed to create file: %s\n", fileName)
	}
	defer f.Close()
	_, err = io.Copy(f, conn)
	if err != nil {
		log.Fatalf("failed to copy: %s\n", err)
	}
}
