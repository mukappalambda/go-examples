package main

import (
	"context"
	"fmt"
	"time"
)

type Item struct {
	Name string
	Age  int
}

func main() {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	items := []Item{
		{
			Name: "alpha",
			Age:  11,
		},
		{
			Name: "beta",
			Age:  22,
		},
		{
			Name: "gamma",
			Age:  33,
		},
	}
	c := Generator(ctx, items...)

	for item := range c {
		fmt.Printf("item: %+v\n", item)
	}
}

func Generator(ctx context.Context, items ...Item) <-chan Item {
	c := make(chan Item)
	go func() {
		defer close(c)
		for _, item := range items {
			select {
			case <-ctx.Done():
				return
			case c <- item:
			}
		}
	}()
	return c
}
