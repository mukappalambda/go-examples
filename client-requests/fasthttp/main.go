package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	client := &fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://example.com")
	req.Header.SetMethod(fasthttp.MethodGet)
	res := fasthttp.AcquireResponse()
	err := client.Do(req, res)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	defer fasthttp.ReleaseResponse(res)
	fmt.Println(string(res.Body()))
}
