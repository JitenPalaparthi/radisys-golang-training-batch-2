package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	//done := make(chan struct{})
	go sender(ch1)
	go sender(ch2)

	go func() {
		for {
			val, ok := <-ch1 // How does it knows the channel is closed.
			// channel returns two values one is val from channel and second one is bool,
			// true is channel is not closed, false channel is closed

			if ok {
				fmt.Println(fmt.Sprint(val) + "--->ch1")
			} else {
				return
			}
		}
		//	done <- struct{}{}
	}()

	// range loop iterates , receives value from channel until a channel is closed
	for val := range ch2 {
		fmt.Println(fmt.Sprint(val) + "---> ch2")
	}

}

func sender(send chan int) {
	i := 1
	for {
		if i > 1000 {
			close(send)
			return
		}
		send <- i
		///.Sleep(time.Microsecond * 1)
		i++
	}
}
