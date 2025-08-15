package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	c := producer(ctx, 5)
	consumer(ctx, c)
}

func producer(ctx context.Context, count int) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for i := range count {
			select {
			case <-ctx.Done():
				return
			case c <- fmt.Sprintf("message-%d", i):
			}
		}
	}()
	return c
}

func consumer(ctx context.Context, c <-chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-c:
			if !ok {
				return
			}
			fmt.Println("received:", v)
		}
	}
}
