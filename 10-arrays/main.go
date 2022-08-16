package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// var v1 int
	// const c1 int
	// var arr [5]int

	var arr1 [5]int = [5]int{10, 12, 13, 15, 16}
	fmt.Println(arr1)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Intn(1000)
	}
	fmt.Println(arr1)
	arr2 := [5]int{10, 4, 9, 10, 15} // short hand declaration
	fmt.Println(arr2)
	//[...] --> this must be evalauated at compile time
	arr3 := [...]int{10, 14, 15, 16, 18, 19, 20, 1, 5, 34, 72, 21} // shorthand declaration
	fmt.Println(arr3)
	fmt.Println("Length of an array:", len(arr3))
	fmt.Println("Capacity of an array:", cap(arr3))
}
