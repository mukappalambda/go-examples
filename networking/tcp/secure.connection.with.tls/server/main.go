package main

import (
	"bufio"
	"crypto/tls"
	"io"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server-key.pem")
	if err != nil {
		log.Fatalf("error loading server cert and key: %s\n", err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	ln, err := tls.Listen("tcp", ":2000", tlsConfig)
	if err != nil {
		log.Fatalf("error creating tls listener on the network: %s\n", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			defer conn.Close()
			reader := bufio.NewReader(conn)
			for {
				s, err := reader.ReadString('\n')
				if err != nil && err == io.EOF {
					break
				}
				log.Printf("client said %q", s)
				if _, err := conn.Write([]byte(s)); err != nil {
					log.Println(err)
					continue
				}
			}
		}()
	}
}
