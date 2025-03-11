package main

import (
	"net/http/httptest"
)

func NewHTTPServer() *httptest.Server {
	return httptest.NewServer(defaultHandler())
}
