package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	numGoroutines := 20
	c := make(chan struct{}, numGoroutines)

	for i := range math.MaxInt32 {
		c <- struct{}{}
		go func(v int) {
			fmt.Println(v)
			time.Sleep(time.Second)
			<-c
		}(i)
	}
}
