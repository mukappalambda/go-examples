package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup

	numGoroutines := 10000

	var cnt int32 = 0
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			time.Sleep(50 * time.Millisecond)
			atomic.AddInt32(&cnt, 1)
			// cnt++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
	fmt.Println(runtime.NumGoroutine())
}
