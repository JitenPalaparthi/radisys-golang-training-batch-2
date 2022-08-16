package main

import (
	"fmt"
	"reflect"
)

// for range loop
// 1- working with arrays and slices --> index, value
// 2- working with maps 			 --> key, value
// 3- working with channels			 --> value
func main() {

	str := "Hello World!"

	for i, v := range str {
		fmt.Println("index:", i, " value:", string(v), "type:", reflect.TypeOf(v))
	}

	for _, v := range str {
		fmt.Println("Val:", string(v))
	}
	for i, _ := range str {
		fmt.Println("index:", i)
	}
	for i := range str {
		fmt.Println("index:", i)
	}

	count := 1
	for range str {
		fmt.Println("How many times this gets printed-->", count)
		count++
	}

}

// 1- find biggest number amoung a,b,c
// 2- reverse a string using for-range loop
