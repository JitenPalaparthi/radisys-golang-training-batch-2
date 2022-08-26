package main

import (
	"fmt"
	"runtime"
)

func main() {
	//runtime.GOMAXPROCS(2)
	go func() {
		i := 0
	loop:
		go fmt.Println(i)
		i++
		if i != 1000000 {
			goto loop
		}
	}()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines those are running:", runtime.NumGoroutine())
	//time.Sleep(time.Millisecond * 5)
	fmt.Println("end of main")
	//runtime.Gosched()
	runtime.Goexit()
}
