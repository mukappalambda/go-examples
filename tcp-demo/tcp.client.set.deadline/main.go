package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	serverDelay   = flag.Duration("server.delay", time.Second, "delay after reading from conn and before writing to conn")
	clientTimeout = flag.Duration("client.timeout", 500*time.Millisecond, "delay after reading from conn and before writing to conn")
)

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("error listening on the tcp network: %s\n", err)
	}
	defer ln.Close()
	addr := ln.Addr().String()
	go func(ln net.Listener) {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go func(conn io.ReadWriteCloser) {
				defer conn.Close()
				reader := bufio.NewReader(conn)
				for {
					line, err := reader.ReadString('\n')
					if err != nil {
						break
					}
					time.Sleep(*serverDelay)
					response := fmt.Sprintf("did you say: %q", line)
					_, err = conn.Write([]byte(response))
					if err != nil {
						continue
					}
				}
			}(conn)
		}
	}(ln)
	fmt.Printf("tcp server running at %s\n", addr)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("error establishing the connection with server: %s\n", err)
	}
	defer conn.Close()
	fmt.Printf("client set up the connection from %s\n", conn.LocalAddr().String())
	if err := conn.SetDeadline(time.Now().Add(*clientTimeout)); err != nil {
		conn.Close()
		log.Fatalf("error setting deadline: %s\n", err)
	}
	fmt.Printf("client set timeout to %s\n", (*clientTimeout).String())
	if _, err := conn.Write([]byte("hi there\n")); err != nil {
		conn.Close()
		log.Fatalf("error writing to server: %s\n", err)
	}
	log.Println("client wrote to server successfully")
	rb := make([]byte, 128)
	if _, err := conn.Read(rb); err != nil {
		conn.Close()
		log.Fatalf("error reading from server: %s\n", err)
	}
}
