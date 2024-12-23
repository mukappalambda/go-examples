package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var connMap = make(map[string]*net.Conn, 0)

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error listening at %d\n", port)
	}
	defer ln.Close()
	fmt.Printf("server listening at %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("error accepting new connection: %s\n", err)
		}
		go HandleConn(conn)
	}
}

func remoteAddr(conn net.Conn) string {
	return conn.RemoteAddr().String()
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	clientAddr := remoteAddr(conn)
	connMap[clientAddr] = &conn
	msg := fmt.Sprintf("[client@%s] joined the room\t[%d guests online]\n", clientAddr, len(connMap))
	log.Print(msg)
	broadcast(msg)
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			delete(connMap, clientAddr)
			log.Printf("error reading from [client@%s]\n", clientAddr)
			msg := fmt.Sprintf("[client@%s] left the room\t[%d guests online]\n", clientAddr, len(connMap))
			log.Print(msg)
			broadcast(msg)
			return
		}
		trimmedLine := strings.TrimSuffix(line, "\n")
		if len(trimmedLine) > 0 {
			msg := fmt.Sprintf("[client@%s]> %s\n", clientAddr, trimmedLine)
			broadcast(msg)
		}
	}
}

func broadcast(msg string) {
	for cAddr, c := range connMap {
		_, err := (*c).Write([]byte(msg))
		if err != nil {
			(*c).Close()
			delete(connMap, cAddr)
			log.Printf("error writing to client: %s\n", cAddr)
			continue
		}
	}
}
