package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	Score int
	Error error
}

func main() {
	c := worker()

	result := <-c

	if err := result.Error; err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Score:", result.Score)
}

func worker() <-chan Result {
	c := make(chan Result)

	go func() {
		defer close(c)
		var result Result
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		v := r1.Intn(10)

		if v < 4 {
			result = Result{Score: -99, Error: errors.New("too low")}
		} else {
			result = Result{Score: v, Error: nil}
		}

		c <- result
	}()

	return c
}
