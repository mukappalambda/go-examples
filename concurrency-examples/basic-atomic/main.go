package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	numGoroutines := 10000

	var cnt int32 = 0
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			atomic.AddInt32(&cnt, 1)
			// cnt++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
	// fmt.Println(runtime.NumGoroutine())
}