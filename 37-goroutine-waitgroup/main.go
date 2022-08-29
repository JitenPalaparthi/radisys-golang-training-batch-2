package main

import (
	"fmt"
	"sync"
)

var num int16 = 1

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1001)
	go func() {
		for i := 1; i <= 1000; i++ {
			go increment(wg)
			//go increment(wg)
		}
		wg.Done()
	}()
	fmt.Println("end of main")
	wg.Wait()
	//runtime.Goexit()
}

func increment(wg *sync.WaitGroup) {
	fmt.Print(num, " ")
	num++
	wg.Done() // when ever wg sees a done, it decrements the counter
}

//waitGroup
// increment
// decrement
// separate go routines.. must be callede 100 times increments
// and 100 times decrement
// use mutex, waitgroup .. at last when you print the value it must be 0
