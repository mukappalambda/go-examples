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
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error connecting to the network: %s\n", err)
	}
	log.Printf("connected to the network successfully")
	defer conn.Close()
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("error returning working directory: %s\n", err)
		conn.Close()
		os.Exit(1)
	}
	fileName := fmt.Sprintf("%s/my_file.txt", dir)
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("error opening file: %s\n", err)
		conn.Close()
		os.Exit(1)
	}
	n, err := io.Copy(conn, f)
	if err != nil {
		log.Printf("failed to send file to tcp connection: %v\n", err)
		f.Close()
		conn.Close()
		os.Exit(1)
	}
	fmt.Printf("sent file content successfully: %d bytes sent\n", n)
}
