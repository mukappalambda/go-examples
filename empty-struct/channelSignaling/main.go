package main

import (
	"fmt"
)

func main() {
	c := make(chan struct{})
	go runner(c)
	fmt.Println("Referee: Ready...")
	fmt.Println("Referee: Go!")
	c <- struct{}{}
	<-c
	fmt.Println("Referee: Congrats!")
	close(c)
}

func runner(c chan struct{}) {
	<-c
	fmt.Println("Runner: (Start to run).")
	fmt.Println("Runner: (Crossed the finish line).")
	c <- struct{}{}
}
