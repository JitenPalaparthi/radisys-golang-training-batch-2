package main

import "fmt"

func main() {
	ch1 := sender()
	ch2 := sender()
	done1 := receiver(ch1)
	done2 := receiver(ch2)
	<-done1
	<-done2

	ch3, ch4, ch5 := multiChanSender()
	done3 := receiver(ch3)
	done4 := receiver(ch4)
	done5 := receiver(ch5)

	<-done3
	<-done4
	<-done5

}

func sender() <-chan int {
	send := make(chan int)
	go innerSender(send)
	return send // a channel has been created and given access outside of func
}

func receiver(receive <-chan int) chan struct{} {
	done := make(chan struct{})
	go func() {
		for val := range receive {
			fmt.Println(val)
		}
		done <- struct{}{}
	}()
	return done
}
func innerSender(send chan int) {
	for i := 1; i <= 100; i++ {
		send <- i
	}
	close(send)
}

func multiChanSender() (<-chan int, <-chan int, <-chan int) {
	send1 := make(chan int)
	send2 := make(chan int)
	send3 := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			send1 <- i
			send2 <- i
			send3 <- i
		}
		close(send1)
		close(send2)
		close(send3)
	}()

	return send1, send2, send3 // a channel has been created and given access outside of func
}

// guidelines using channels
// 1- if same goroutine which is a sender is called multiple times, then if the channel is closed
// 	  in one of the goroutines, then the other goroutine does not know about channel status and then try to send values.
// To avoid issue in point 1 , use different channels for different goroutines particularly
// 	they are senders.
// 2- Never use go routune inside a loop. Use as goroutine outside of the loop.
