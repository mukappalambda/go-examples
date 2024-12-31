package main

import (
	"compress/gzip"
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
	b = []byte(fmt.Sprintf("<wrapped by middleware>%s<wrapped by middleware>", string(b)))
	return grw.zw.Write(b)
}

func main() {
	http.HandleFunc("GET /uncompressed", handleData())
	http.HandleFunc("GET /compressed", compressionMiddleware(handleData()))
	ts := httptest.NewServer(nil)
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL + "/uncompressed")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	res, err = client.Get(ts.URL + "/compressed")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}
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
