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
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	wg.Add(1)
	go func (ctx context.Context) {
		defer wg.Done()
		worker(ctx)
	}(ctx)
	wg.Wait()
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func worker(ctx context.Context) {
	fmt.Println("worker is up!")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop working.")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("worker is working...")
		}
	}
}