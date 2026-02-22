package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	numGoroutines := 10
	wg.Add(numGoroutines)
	for i := range numGoroutines {
		go func(ctx context.Context, id int) {
			defer wg.Done()
			worker(ctx, id)
		}(ctx, i)
	}
	wg.Wait()
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func worker(ctx context.Context, id int) {
	fmt.Println("worker", id, "is up!")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "stops working.")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("worker", id, "is working...")
		}
	}
}
