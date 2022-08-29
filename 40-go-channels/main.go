package main

import "fmt"

func main() {
	ch := make(chan string) // bi-directional channel
	done := make(chan struct{})
	go sender(ch, "sender-1")
	go sender(ch, "sender-2")
	go receiver(ch, done, "receiver-1")
	go receiver(ch, done, "receiver-2")
	<-done
	<-done
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(<-ch)
	// }
}

func sender(send chan<- string, str string) { // send only channel
	for i := 1; i <= 10; i++ {
		send <- "Hello World!-->" + fmt.Sprint(i) + "-->" + str
	}
}

func receiver(receive <-chan string, done chan<- struct{}, str string) { // receive only channel
	for i := 1; i <= 10; i++ {
		msg := <-receive
		fmt.Println(str + "-->" + msg)
	}
	done <- struct{}{}
}
