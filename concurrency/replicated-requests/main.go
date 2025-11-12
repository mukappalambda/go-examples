package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"time"
)

var n = flag.Int("n", 1, "number of replicated requests")

func main() {
	flag.Parse()
	start := time.Now()
	out := make(chan string)
	query := "test"
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(*n)
	for range *n {
		go func() {
			defer wg.Done()
			result, err := fakeAPI(ctx, query)
			if err != nil {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- result:
			}
		}()
	}
	first := <-out
	_ = first
	cancel()
	elapsed := time.Since(start)
	fmt.Println("elapsed:", elapsed)
	wg.Wait()
	fmt.Println("goroutine counts:", runtime.NumGoroutine())
}

func fakeAPI(ctx context.Context, query string) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return "", err
	}
	random := n.Int64()
	timer := time.NewTimer(time.Duration(random) * time.Millisecond)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-timer.C:
		return fmt.Sprintf("result of %s", query), nil
	}
}
