package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

const caFile = "root_ca.crt"

func main() {
	file, err := os.Open(caFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	caCert, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %s\n", err)
		file.Close()
		return
	}
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(caCert)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			RootCAs:    certpool,
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	url := "https://localhost:9443/data"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create request: %s\n", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get request: %s\n", err)
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read response body: %s\n", err)
		return
	}
	fmt.Println(string(content))
}
