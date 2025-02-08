package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	addr := "127.0.0.1:8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error listening on the network: %s\n", err)
	}
	defer ln.Close()
	log.Printf("server listening at: %s\n", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("error accepting the connection: %s\n", err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()
			reader := bufio.NewReader(conn)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					log.Printf("error reading from connection: %s\n", err)
					break
				}
				text := strings.TrimSuffix(line, "\n")
				var data map[string]interface{}
				err = json.Unmarshal([]byte(text), &data)
				if err != nil {
					log.Printf("error unmarshaling input: %s\n", err)
					continue
				}
				data["received_at"] = time.Now().Format(time.DateTime)
				b, err := json.Marshal(&data)
				if err != nil {
					log.Printf("error marshaling response: %s\n", err)
					continue
				}
				b = append(b, '\n')
				_, err = conn.Write(b)
				if err != nil {
					log.Println(err)
					conn.Close()
					break
				}
			}
		}(conn)
	}
}
