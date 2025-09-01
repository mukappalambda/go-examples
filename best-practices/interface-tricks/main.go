package main

import (
	"flag"
	"fmt"
)

type Ingester interface {
	Ingest([]float64) []float64
}

type IngesterFunc func([]float64) []float64

func (i IngesterFunc) Ingest(data []float64) []float64 {
	return i(data)
}

func addOne(data []float64) []float64 {
	out := make([]float64, len(data))
	for k, v := range data {
		out[k] = v + 1
	}
	return out
}

func ScaledTransform(i Ingester, scale float64) Ingester {
	return IngesterFunc(func(data []float64) []float64 {
		out := make([]float64, len(data))
		for k, v := range i.Ingest(data) {
			out[k] = scale * v
		}
		return out
	})
}

var scale = flag.Float64("scale", 1.23, "scale coefficient")

func main() {
	flag.Parse()
	data := make([]float64, 5)

	fmt.Printf("initial data: %v\n", data)
	inputIngester := IngesterFunc(addOne)
	fmt.Printf("data by addOne: %v\n", inputIngester.Ingest(data))
	outputIngester := ScaledTransform(inputIngester, *scale)
	fmt.Printf("data by addOne then scaled by %f: %v\n", *scale, outputIngester.Ingest(data))
}
