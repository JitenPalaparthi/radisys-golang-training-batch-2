package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	ch <- 200

	fmt.Println("Len:", len(ch), "Cap:", cap(ch))
	fmt.Println(<-ch, <-ch)
	fmt.Println("Len:", len(ch), "Cap:", cap(ch))
	ch <- 300
	ch1 := sender()
	fmt.Println("Len:", len(ch1), "Cap:", cap(ch1))

	done1 := receiver(ch1)
	done2 := receiver(ch1)

	<-done1
	<-done2

}

func sender() <-chan int {
	send := make(chan int, 2)
	go func() {
		for i := 1; i <= 100; i++ {
			send <- i
		}
		close(send)
	}()
	return send
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
