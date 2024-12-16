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
		log.Fatal(err)
	}
	defer conn.Close()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fileName := fmt.Sprintf("%s/my_file.txt", dir)
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s\n", err)
	}
	_, err = io.Copy(conn, f)
	if err != nil {
		log.Fatalf("failed to send file to tcp connection: %v\n", err)
	}
	fmt.Println("sent file content successfully.")
}
