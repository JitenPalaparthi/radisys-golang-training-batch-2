package main

import (
	"fmt"
	"runtime"
	"vendors/mb"
)

var (
	CONN []string = []string{"localhost:29092"}
)

func main() {
	fmt.Println("Subscriber started...")
	var chMsg <-chan []byte
	chMsg = mb.Subscribe(CONN, "vendor.created")
	for msg := range chMsg {
		fmt.Println(string(msg))
	}
	runtime.Goexit()
}
