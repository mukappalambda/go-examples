package server

import (
	"net/http"
	"time"
)

type serverOpts struct {
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
	writeTimeout      time.Duration
	handler           http.Handler
}

type Opt func(s *serverOpts)

func WithDefaultReadTimeout(d time.Duration) Opt {
	return func(s *serverOpts) {
		s.readTimeout = d
	}
}

func WithDefaultReadHeaderTimeout(d time.Duration) Opt {
	return func(s *serverOpts) {
		s.readHeaderTimeout = d
	}
}

func WithDefaultWriteTimeout(d time.Duration) Opt {
	return func(s *serverOpts) {
		s.writeTimeout = d
	}
}

func WithDefaultHandler(handler http.Handler) Opt {
	return func(s *serverOpts) {
		s.handler = handler
	}
}
