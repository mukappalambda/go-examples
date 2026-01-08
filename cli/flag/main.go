package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type HyperParam struct {
	Epochs    uint
	BatchSize uint
	Optimizer string
	Criterion string
	Dropout   float64
}

func newHyperParam(epochs, batchSize uint, optimizer, criterion string, dropout float64) HyperParam {
	return HyperParam{
		Epochs:    epochs,
		BatchSize: batchSize,
		Optimizer: optimizer,
		Criterion: criterion,
		Dropout:   dropout,
	}
}

var (
	batchSize = flag.Uint("bs", 32, "batch size")
	criterion = flag.String("criterion", "mse", "criterion")
	dropout   = flag.Float64("dropout", 0.2, "dropout probability")
	epochs    = flag.Uint("epochs", 10, "epochs")
	optimizer = flag.String("optimizer", "adam", "optimizer")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	if *criterion != "mse" && *criterion != "mae" {
		return fmt.Errorf("currently supported criteria are 'mae' or 'mse'")
	}

	if *dropout > 1 {
		return fmt.Errorf("dropout should be in the closed interval [0, 1]")
	}

	if *epochs < 1 {
		return fmt.Errorf("epochs flag should be greater than 0")
	}
	hp := newHyperParam(*epochs, *batchSize, *optimizer, *criterion, *dropout)

	fmt.Printf("%+v\n", hp)
	return nil
}
