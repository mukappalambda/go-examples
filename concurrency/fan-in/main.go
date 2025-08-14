package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	c1 := producer(ctx, "alpha", 2)
	c2 := producer(ctx, "beta", 3)
	c3 := producer(ctx, "gamma", 5)
	cs := []<-chan string{c1, c2, c3}
	c := fanIn(ctx, cs...)
	for v := range c {
		fmt.Println("item:", v)
	}
}

func producer(ctx context.Context, id string, count int) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for i := range count {
			select {
			case <-ctx.Done():
				return
			case c <- fmt.Sprintf("product-%d-producer-%s", i, id):
			}
		}
	}()
	return c
}

func fanIn(ctx context.Context, cs ...<-chan string) <-chan string {
	out := make(chan string)
	if len(cs) == 0 {
		close(out)
		return out
	}
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, ci := range cs {
		go func() {
			defer wg.Done()
			for v := range ci {
				select {
				case <-ctx.Done():
					return
				case out <- v:
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
