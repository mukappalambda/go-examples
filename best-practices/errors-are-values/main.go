package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type errWriter struct {
	client *http.Client
	err    error
}

func NewErrWriter() *errWriter {
	ew := &errWriter{
		client: &http.Client{},
		err:    nil,
	}
	return ew
}

func (ew *errWriter) NewRequest(method string, url string, body io.Reader) *http.Request {
	return ew.NewRequestWithContext(context.Background(), method, url, body)
}

func (ew *errWriter) NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		// This error check may be redundant as it is the possible error that errWriter would encounter
		if ew.err == nil {
			ew.err = err
		}
		return nil
	}
	return req
}

func (ew *errWriter) GetResponseBody(req *http.Request) io.ReadCloser {
	if req == nil {
		// Only want to track the first error, hence the if statement
		if ew.err == nil {
			ew.err = fmt.Errorf("nil http.Request")
		}
		return nil
	}
	resp, err := ew.client.Do(req)
	if err != nil {
		if ew.err == nil {
			ew.err = err
		}
		return nil
	}
	return resp.Body
}

func (ew *errWriter) ReadAll(r io.Reader) []byte {
	if r == nil {
		if ew.err == nil {
			ew.err = fmt.Errorf("nil io.Reader")
		}
		return nil
	}
	data, err := io.ReadAll(r)
	if err != nil {
		if ew.err == nil {
			ew.err = err
		}
		return nil
	}
	return data
}

var url = flag.String("url", "", "URL")

func main() {
	flag.Parse()
	checkURL(*url)

	ew := NewErrWriter()
	method := http.MethodGet
	// No need to check every error after every operation
	req := ew.NewRequest(method, *url, nil)
	respBody := ew.GetResponseBody(req)
	content := ew.ReadAll(respBody)
	// Check the error at the last step
	if ew.err != nil {
		log.Fatal(ew.err)
	}
	// Close the resource
	defer respBody.Close()
	fmt.Printf("data: %s", string(content))
}

// checkURL with exit with return code 1 if url is invalid
func checkURL(url string) {
	if len(url) == 0 {
		fmt.Println("forgot to pass flag -url?")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
