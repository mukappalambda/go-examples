package main

import (
	"fmt"

	"github.com/cpmech/gosl/fun/fftw"
)

type Data struct {
	Value []complex128
}

func (d *Data) ForwardFFT() {
	forwardFft := fftw.NewPlan1d(d.Value, false, false)
	forwardFft.Execute()
	forwardFft.Free()
}

func (d *Data) BackwardFFT() {
	backwardFft := fftw.NewPlan1d(d.Value, true, false)
	backwardFft.Execute()
	backwardFft.Free()
}

func main() {
	data := Data{
		Value: []complex128{1, 2, 3, 4},
	}
	data.ForwardFFT()
	fmt.Printf("forward fft: %v\n", data.Value)
	data.BackwardFFT()
	fmt.Printf("backward fft: %v\n", data.Value)
}
