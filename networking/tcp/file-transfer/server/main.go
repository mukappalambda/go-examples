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
		log.Fatalf("error listening at %d\n", port)
	}
	fmt.Printf("server is listening at %s\n", ln.Addr().String())
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("error accepting a new connection: %s\n", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	log.Printf("client@%s is connected\n", remoteAddr)
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("error returning working directory: %s\n", err)
		conn.Close()
		return
	}
	fileName := fmt.Sprintf("%s/tmp_file.txt", dir)
	f, err := os.Create(fileName)
	if err != nil {
		log.Printf("error creating a new file: %s\n", fileName)
		conn.Close()
		return
	}
	defer f.Close()
	n, err := io.Copy(f, conn)
	if err != nil {
		log.Printf("error copying to another connection: %s\n", err)
		f.Close()
		conn.Close()
	}
	log.Printf("%s sent %d bytes successfully\n", remoteAddr, n)
}
