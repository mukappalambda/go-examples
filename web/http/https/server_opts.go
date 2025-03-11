package main

import (
	"net/http"
	"time"
)

type serverOpt struct {
	handler           http.Handler
	readHeaderTimeout time.Duration
	readTimeout       time.Duration
}

type Opt func(opt *serverOpt)

func WithDefaultHandler(h http.Handler) Opt {
	return func(opt *serverOpt) {
		opt.handler = h
	}
}

func WithDefaultReadHeaderTimeout(d time.Duration) Opt {
	return func(opt *serverOpt) {
		opt.readHeaderTimeout = d
	}
}

func WithDefaultReadTimeout(d time.Duration) Opt {
	return func(opt *serverOpt) {
		opt.readTimeout = d
	}
}
