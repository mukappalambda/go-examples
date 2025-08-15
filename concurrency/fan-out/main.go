package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"strings"
	"sync"
	"time"
)

var (
	n = flag.Int("n", 50, "number of items")
	j = flag.Int("j", 1, "number of workers")
)

func main() {
	flag.Parse()
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	start := time.Now()
	in := producer(ctx, *n)
	out := fanOut(ctx, in, *j, toUpperFunc)
	for v := range out {
		_ = v
	}
	elapsed := time.Since(start)
	fmt.Println("elapsed", elapsed)
}

func producer(ctx context.Context, cnt int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := range cnt {
			select {
			case <-ctx.Done():
				return
			case out <- fmt.Sprintf("item-%d", i):
			}
		}
	}()
	return out
}

func fanOut(ctx context.Context, in <-chan string, workerCount int, workerFunc func(string) string) <-chan string {
	out := make(chan string)
	if workerCount < 0 {
		close(out)
		return out
	}
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for range workerCount {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					out <- workerFunc(v)
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func toUpperFunc(msg string) string {
	random := rand.IntN(10)
	time.Sleep(time.Duration(random) * time.Millisecond)
	return strings.ToUpper(msg)
}
