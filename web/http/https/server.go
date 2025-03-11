package main

import (
	"net/http"
	"time"
)

func NewServer(addr string, opt ...Opt) *http.Server {
	var sopts serverOpt
	for _, o := range opt {
		o(&sopts)
	}
	if sopts.readHeaderTimeout == 0 {
		sopts.readHeaderTimeout = 5 * time.Second
	}
	if sopts.readTimeout == 0 {
		sopts.readTimeout = 5 * time.Second
	}
	return &http.Server{
		Addr:              addr,
		Handler:           sopts.handler,
		ReadHeaderTimeout: sopts.readHeaderTimeout,
		ReadTimeout:       sopts.readTimeout,
	}
}
