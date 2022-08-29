package main

import "fmt"

// 1- To create a go channel use chan keyword.
// 2- To instantiate use make function: make(chan int)
// 3- There are two types of gochannels buffered and unbuffered
// 4- While creating channel no size if given then it is a unbuffered channel
// 5- The type of the chan is the kind of value that channel can send or receive
// 6- For unbuffered channels,the sender is blocked until receiver receives value
// 7- For unbuffered channels, the receiver is blocked until sender sends the value
// 8- Channels to be used for communication between goroutines
// 9- To send a value to channel ch <- 100
// 10- To receive a value from a chennal val:= <-ch
// 11- not to use same channel to send or receive in the same goroutine

func main() {
	var ch chan int     //declaring a channel; ch is nil as not instantiated
	ch = make(chan int) // ch is instantiated but unbuffered
	// ch <- 100           // who is sender? Sender is main goroutine.
	// // Who is blocked? main is blocked
	// val := <-ch
	go func() {
		ch <- 100
	}()
	val := <-ch
	fmt.Println(val)
	// val := 0
	// go func() {
	// 	val = <-ch
	// 	fmt.Println(val)
	// }()
	// ch <- 100
	// time.Sleep(time.Second * 1)
}
