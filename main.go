package main

import (
	"fmt"
	"time"
)

// Only Send Channel
func countToTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i // send i to channel
	}
	close(c) // avoid deadlock
}

// Only Receive Channel
func receive(c <-chan int) {
	for {
		a, ok := <-c
		if !ok {
			fmt.Println("Done")
			break
		}
		fmt.Printf("received %d\n", a)
	}
}

func main() {
	c := make(chan int)
	go countToTen(c)
	receive(c)
}
