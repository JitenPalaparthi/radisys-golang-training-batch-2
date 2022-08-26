package main

import (
	"fmt"
	"time"
)

// 1- main is also a go routine
// 2- to create a go routine use a keyword called go
// 3- by default no go routine waits for other go routine to complete its execution
// 4- by default the order of execution cannot be determined

func main() {
	go func() {
		fmt.Println("Hello World-Through Anonymous Func")
	}()
	go greet()
	fmt.Println("end of main")
	time.Sleep(time.Millisecond * 1)
}

func greet() {
	// go fmt.Println("Hello World!")
	go fmt.Print("Hello")
	go fmt.Print(" World!")
	go fmt.Println("----------")
}

// simple go routine as an anonymous func
// take a loop of 10000 elements
// print them as a separate go routine. also execute the loop as a separate go routine..
