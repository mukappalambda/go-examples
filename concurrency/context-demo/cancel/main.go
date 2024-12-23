package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
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
