package main

import (
	"fmt"
	"sync"
)

var num int16 = 1

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go increment(wg)
		//go increment(wg)
	}
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
