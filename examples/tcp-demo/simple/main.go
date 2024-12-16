package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	port := 8080
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server is listening at %d\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("client from %q is connected.\n", remoteAddr)
	reader := bufio.NewReader(conn)
	for {
		buf, err := reader.ReadString('\n')
		buf = strings.Trim(buf, "\n")
		if err != nil {
			log.Fatal(err)
		}
		var ts string
		ts = time.Now().Local().Format(time.DateTime)
		fmt.Printf("[client@%s] %s > %s\n", remoteAddr, ts, buf)
		ts = time.Now().Local().Format(time.DateTime)
		msg := fmt.Sprintf("[server@%s] %s > haha you said: %q?\n", localAddr, ts, buf)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
	}
}
