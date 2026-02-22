package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	shutdown := make(chan struct{})

	wg.Go(func() {

		for {
			select {
			case <-shutdown:
				fmt.Println("Cleaning resources...", time.Now().UTC().Truncate(time.Millisecond))
				time.Sleep(time.Second)
				fmt.Println("Resources have been cleaned. Bye.", time.Now().UTC().Truncate(time.Millisecond))
				return
			default:
				myFunc()
			}
		}
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Shutting shutdown...")
		close(shutdown)
	}()

	wg.Wait()
}

// Emulate an expensive computation
func myFunc() {
	time.Sleep(2 * time.Second)
	fmt.Println("Task completed", time.Now().UTC().Truncate(time.Millisecond))
}
