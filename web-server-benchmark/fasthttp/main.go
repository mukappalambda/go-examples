package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var port = flag.Int("port", 8080, "server port")

func main() {
	flag.Parse()

	fmt.Printf("server listening at %d\n", *port)
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", *port), indexHandler); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello from the fasthttp server.")
}
