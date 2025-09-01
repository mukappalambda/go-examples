package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type App struct {
	s *http.Server
}

var (
	defaultAddr              = ":8080"
	defaultReadTimeout       = 500 * time.Millisecond
	defaultReadHeaderTimeout = 500 * time.Millisecond

	defaultAppOption = &appOption{
		addr:              defaultAddr,
		handler:           nil,
		readTimeout:       defaultReadTimeout,
		readHeaderTimeout: defaultReadHeaderTimeout,
	}
)

var (
	addr              = flag.String("addr", defaultAddr, "server addr")
	readTimeout       = flag.Duration("read-timeout", defaultReadTimeout, "read timeout")
	readHeaderTimeout = flag.Duration("read-header-timeout", defaultReadHeaderTimeout, "read header timeout")
)

func main() {
	flag.Parse()
	app := NewApp(WithAddrOption(*addr), WithReadTimeout(*readTimeout), WithReadHeaderTimeout(*readHeaderTimeout))
	fmt.Printf("server listening at %v\n", app.Addr())
	if err := app.Run(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type appOption struct {
	addr              string
	handler           http.Handler
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
}

type AppOption interface {
	apply(*appOption)
}

type Addr struct {
	string
}

func (a Addr) apply(o *appOption) {
	o.addr = a.string
}

func WithAddrOption(addr string) AppOption {
	return Addr{string: addr}
}

type ReadTimeout struct {
	d time.Duration
}

func (rt ReadTimeout) apply(o *appOption) {
	o.readTimeout = rt.d
}

func WithReadTimeout(d time.Duration) AppOption {
	return ReadTimeout{d: d}
}

type ReadHeaderTimeout struct {
	d time.Duration
}

func (rht *ReadHeaderTimeout) apply(o *appOption) {
	o.readHeaderTimeout = rht.d
}

func WithReadHeaderTimeout(d time.Duration) AppOption {
	return &ReadHeaderTimeout{d: d}
}

func NewApp(opt ...AppOption) *App {
	opts := defaultAppOption
	for _, o := range opt {
		o.apply(opts)
	}
	s := &http.Server{
		Addr:              opts.addr,
		Handler:           opts.handler,
		ReadTimeout:       opts.readTimeout,
		ReadHeaderTimeout: opts.readHeaderTimeout,
	}
	a := &App{
		s: s,
	}
	return a
}

func (a *App) Addr() string {
	return a.s.Addr
}

func (a *App) Run() error {
	return a.s.ListenAndServe()
}
