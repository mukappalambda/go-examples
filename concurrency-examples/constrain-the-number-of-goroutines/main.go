package main

import (
	"fmt"
	"math"
	"time"
)

func main()  {
	numGoroutines := 20
	c := make(chan struct{}, numGoroutines)

	for i := 0; i < math.MaxInt32; i++ {
		c <-struct{}{}
		go func (v int) {
			fmt.Println(v)
			time.Sleep(time.Second)
			<-c
		}(i)
	}
	
}