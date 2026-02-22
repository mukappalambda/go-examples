package main

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	zw *gzip.Writer
}

func (grw *gzipResponseWriter) Write(b []byte) (int, error) {
	b = fmt.Appendf(nil, "<wrapped by middleware>%s<wrapped by middleware>", string(b))
	return grw.zw.Write(b)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	http.HandleFunc("GET /uncompressed", handleData())
	http.HandleFunc("GET /compressed", compressionMiddleware(handleData()))
	ts := httptest.NewServer(nil)
	defer ts.Close()

	client := ts.Client()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/uncompressed", nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		return err
	}

	fmt.Println()
	req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL+"/compressed", nil)
	if err != nil {
		return err
	}

	res, err = client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		return err
	}
	return nil
}

func compressionMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		zw := gzip.NewWriter(w)
		defer zw.Close()
		grw := &gzipResponseWriter{ResponseWriter: w, zw: zw}
		h.ServeHTTP(grw, r)
	}
}

func handleData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, "data"); err != nil {
			errMsg := "error writing to response writer"
			log.Println(errMsg, err)
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
	}
}
