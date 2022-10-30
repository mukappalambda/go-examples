package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	done := make(chan struct{})

	c := sender(done)

	wg.Add(1)
	go func(ch <-chan string) {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
			time.Sleep(time.Second)
		}
	}(c)
	wg.Wait()
	fmt.Println("Closing the channel to cancel the sender.")
	close(done)
	fmt.Println(<-c)

}

func sender(done <-chan struct{}) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for {
			select {
			case <-done:
				c <- "sender: bye!"
				return
			case c <- "sender pings!":
			}
		}
	}()
	return c
}
