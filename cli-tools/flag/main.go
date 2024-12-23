package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type HyperParam struct {
	Epochs    int
	BatchSize int
	Optimizer string
	Criterion string
	Dropout   float64
}

var (
	batchSize = flag.Int("bs", 32, "batch size")
	criterion = flag.String("criterion", "mse", "criterion")
	dropout   = flag.Float64("dropout", 0.2, "dropout probability")
	epochs    = flag.Int("epochs", 10, "epochs")
	optimizer = flag.String("optimizer", "adam", "optimizer")
)

func main() {
	flag.Parse()

	if *criterion != "mse" && *criterion != "mae" {
		fmt.Println("currently supported criteria are 'mae' or 'mse'")
		os.Exit(1)
	}

	if *dropout < 0 || *dropout > 1 {
		fmt.Println("dropout should be in the closed interval [0, 1]")
		os.Exit(1)
	}

	if *epochs < 1 {
		log.Println("epochs flag should be greater than 0")
		os.Exit(1)
	}

	hyperParam := HyperParam{
		Epochs:    *epochs,
		BatchSize: *batchSize,
		Optimizer: *optimizer,
		Criterion: *criterion,
		Dropout:   *dropout,
	}

	fmt.Printf("%+v\n", hyperParam)
	// {Epochs:10 BatchSize:32 Optimizer:adam Criterion:mse Dropout:0.2}
}
