package handlers

import (
	"net/http"
	"time"
)

func DefaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /data", handleData())
	mux.HandleFunc("GET /hello", handleGreet("hello"))
	mux.HandleFunc("GET /hiii", handleGreet("hiii"))
	mux.HandleFunc("GET /slow", handleSlow(3*time.Second))
	return mux
}
