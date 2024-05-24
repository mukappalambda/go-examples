package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := producer("alex", 1)
	c2 := producer("bob", 2)
	c3 := producer("mark", 3)
	c := fanIn(c1, c2, c3)

	for v := range c {
		fmt.Println(v)
	}
}

func producer(name string, count int) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for i := 0; i < count; i++ {
			time.Sleep(time.Second)
			c <- name
		}
	}()

	return c
}

func fanIn(cs ...<-chan string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, ci := range cs {
			for v := range ci {
				c <- v
			}
		}
	}()

	return c
}
