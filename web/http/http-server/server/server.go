package server

import (
	"context"
	"net/http"

	"github.com/mukappalambda/go-examples/web/http/http-server/handlers"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string, opts ...Opt) *Server {
	var sopts serverOpts
	for _, o := range opts {
		o(&sopts)
	}
	if sopts.readTimeout == 0 {
		sopts.readTimeout = defaultReadTimeout
	}
	if sopts.readHeaderTimeout == 0 {
		sopts.readTimeout = defaultReadHeaderTimeout
	}
	if sopts.writeTimeout == 0 {
		sopts.writeTimeout = defaultWriteTimeout
	}
	if sopts.handler == nil {
		sopts.handler = handlers.DefaultMux()
	}
	srv := &http.Server{
		Addr:              addr,
		ReadTimeout:       sopts.readTimeout,
		ReadHeaderTimeout: sopts.readHeaderTimeout,
		WriteTimeout:      sopts.writeTimeout,
		Handler:           sopts.handler,
	}
	return &Server{
		httpServer: srv,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
