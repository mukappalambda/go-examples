package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	port := 8080
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("error connecting to the network: %w", err)
	}
	log.Printf("connected to the network successfully")
	defer conn.Close()
	dir, err := os.Getwd()
	if err != nil {
		conn.Close()
		return fmt.Errorf("error returning working directory: %w", err)
	}
	fileName := fmt.Sprintf("%s/my_file.txt", dir)
	f, err := os.Open(fileName)
	if err != nil {
		conn.Close()
		return fmt.Errorf("error opening file: %w", err)
	}
	n, err := io.Copy(conn, f)
	if err != nil {
		f.Close()
		conn.Close()
		return fmt.Errorf("failed to send file to tcp connection: %w", err)
	}
	fmt.Printf("sent file content successfully: %d bytes sent\n", n)
	return nil
}
