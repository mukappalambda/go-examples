package main

import (
	"io"

	"github.com/valyala/fasthttp"
)

func main() {
	fasthttp.ListenAndServe(":8081", indexHandler)
}

func indexHandler(ctx *fasthttp.RequestCtx) {
	io.WriteString(ctx, "Hello from the fasthttp server.")
}
