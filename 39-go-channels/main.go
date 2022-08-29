package main

import (
	"fmt"
)

// channel of empty struct is used as a trigger that notifies
func main() {
	ch := make(chan int)        // this is unbuffered channel
	done := make(chan struct{}) // this is also unbuffered channel.channel with empty struct
	go func() {
		ch <- 100
	}()
	go func() {
		fmt.Println(<-ch)
		done <- struct{}{} // sending to a channel.
		// What are you sending? Nothing. Still there is a trigger
	}()
	<-done // receiving from a channel; What do you receive? Receive nothing but this is triggered.
	//sum(10, 20) // This is valid,but of no use
}

func sum(a, b int) int {
	return a + b
}

// increment and decrement multiple go routines
// instead of waitgroup use done channel
