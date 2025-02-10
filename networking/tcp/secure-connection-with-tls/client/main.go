package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"os"
)

func main() {
	pemCertsFile, err := os.Open("certs/ca.pem")
	if err != nil {
		log.Fatalf("error opening cert file: %s\n", err)
	}
	defer pemCertsFile.Close()
	pemCerts, err := io.ReadAll(pemCertsFile)
	if err != nil {
		log.Fatalf("error reading certs: %s\n", err)
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(pemCerts)
	if !ok {
		log.Fatal("error parsing root certs")
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:2000", &tls.Config{RootCAs: roots})
	if err != nil {
		log.Fatalf("error connecting to server: %s\n", err)
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("hello from client\n")); err != nil {
		log.Fatalf("error writing to the connection: %s\n", err)
	}
}
