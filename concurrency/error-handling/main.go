package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

type Result struct {
	Score int
	Error error
}

var bound int

func main() {
	c := worker()

	result := <-c
	if err := result.Error; err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Printf("score: %d\n", result.Score)
}

func worker() <-chan Result {
	c := make(chan Result)
	bound = 5
	scoreFn := func(i int) int {
		return 2 * i
	}

	go func() {
		defer close(c)
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		v := int(n.Int64())
		score := scoreFn(v)
		var err error
		if v < bound {
			err = fmt.Errorf("sample is lower than %d: %v", bound, v)
		}
		result := Result{
			Score: score,
			Error: err,
		}

		c <- result
	}()
	return c
}
