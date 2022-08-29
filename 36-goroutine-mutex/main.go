package main

import (
	"fmt"
	"runtime"
	"sync"
)

var num int16 = 1

func main() {
	mu := new(sync.Mutex)
	for i := 1; i <= 1000; i++ {
		go increment(mu)
		go increment(mu)
	}
	runtime.Goexit()
}

func increment(mu *sync.Mutex) {
	mu.Lock()
	fmt.Print(num, " ")
	num++
	mu.Unlock()
}
