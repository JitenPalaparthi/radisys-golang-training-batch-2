package main

import (
	"fmt"
	"runtime"
)

func main() {
	//runtime.GOMAXPROCS(2)
	go func() {
		for i := 1; ; i++ {
			go fmt.Println(i)
		}

	}()

	fmt.Println("CPUs:", runtime.NumCPU())

	fmt.Println("Go routines those are running:", runtime.NumGoroutine())
	//time.Sleep(time.Millisecond * 5)
	fmt.Println("end of main")
	//runtime.Gosched()
	runtime.Goexit()
}
