package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	trace := &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			log.Println("DNS Start")
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			log.Println("DNS Done")
		},
		GotConn: func(ci httptrace.GotConnInfo) {
			log.Printf("Connection got: %+v\n", ci)
		},
		ConnectStart: func(network, addr string) {
			log.Println("Connect Start:", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			log.Println("Connect Done:", network, addr)
		},
	}
	req, err := http.NewRequestWithContext(httptrace.WithClientTrace(context.Background(), trace), http.MethodGet, "https://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
}
