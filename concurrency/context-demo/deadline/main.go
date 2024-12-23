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
	deadline := time.Now().Add(3 * time.Second)
	fmt.Println("deadline:", deadline)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		worker(ctx)
	}(ctx)

	wg.Wait()
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func worker(ctx context.Context) {
	fmt.Println("worker is up!", time.Now())
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			fmt.Println("Stop working.", time.Now())
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("worker is working...", time.Now())
		}
	}
}
