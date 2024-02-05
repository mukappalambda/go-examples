package main

import (
	"fmt"
	"log"
	"math/rand"

	"gonum.org/v1/gonum"
	"gonum.org/v1/gonum/mat"
)

func main() {
	fmt.Println(gonum.Version())

	nRows := 2
	nCols := 3
	data := make([]float64, nRows*nCols)

	for i := range data {
		data[i] = rand.NormFloat64()
	}

	A := mat.NewDense(nRows, nCols, data)
	fmt.Println("A:\n", mat.Formatted(A, mat.Squeeze()))
	r, c := A.Dims()
	fmt.Printf("nRows: %d, nCols: %d\n", r, c)

	if !mat.Equal(A, A.T().T()) {
		log.Fatal("You should not see this.")
		return
	}

	var svd mat.SVD
	ok := svd.Factorize(A, mat.SVDFull)

	if !ok {
		log.Fatal("You should not see this.")
	}

	var U mat.Dense
	sigmas := make([]float64, min(nRows, nCols))
	var V mat.Dense

	svd.UTo(&U)
	svd.Values(sigmas)
	svd.VTo(&V)
	fmt.Println("U:\n", mat.Formatted(&U, mat.Squeeze()))
	fmt.Println("V:\n", mat.Formatted(&V, mat.Squeeze()))
	fmt.Println("sigmas:\n", sigmas)

	var eye1 mat.Dense
	var eye2 mat.Dense
	eye1.Mul(&U, U.T())
	eye2.Mul(&V, V.T())

	fmt.Println("eye1:\n", mat.Formatted(&eye1, mat.Squeeze()))
	fmt.Println("eye2:\n", mat.Formatted(&eye2, mat.Squeeze()))

	eps := 1e-9
	diagElems1 := make([]float64, nRows)
	diagElems2 := make([]float64, nCols)

	for i := range diagElems1 {
		diagElems1[i] = 1
	}

	for i := range diagElems2 {
		diagElems2[i] = 1
	}

	if !mat.EqualApprox(&eye1, mat.NewDiagDense(nRows, diagElems1), eps) {
		log.Fatal("You should not see this.")
		return
	}

	if !mat.EqualApprox(&eye2, mat.NewDiagDense(nCols, diagElems2), eps) {
		log.Fatal("You should not see this.")
		return
	}
}
