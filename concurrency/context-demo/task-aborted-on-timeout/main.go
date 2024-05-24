package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

var (
	dt      = flag.Duration("dt", 10*time.Millisecond, "duration")
	timeout = flag.Duration("timeout", 5*time.Millisecond, "timeout")
)

func main() {
	flag.Parse()
	fmt.Printf("Task execution time: %v; timeout: %v\n", dt, timeout)

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		myFunc(ctx, *dt)
	}()
	wg.Wait()
}

func myFunc(ctx context.Context, d time.Duration) {
	select {
	case <-time.After(d):
		fmt.Println("Task completed.")
	case <-ctx.Done():
		fmt.Printf("Timeout is reached. Task aborted.")
	}
}
