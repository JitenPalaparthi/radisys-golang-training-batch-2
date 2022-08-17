package main

import (
	"fmt"
	"math/rand"
	"reflect"
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
	arr4 := [5]int{}                 // deep copy, value copy
	arr4 = arr2
	arr2[0] = 100
	fmt.Println("Type of arr2:", reflect.TypeOf(arr2))
	fmt.Println(arr2)
	fmt.Println("Type of arr4:", reflect.TypeOf(arr4))
	fmt.Println(arr4)
	//[...] --> this must be evalauated at compile time
	arr3 := [...]int{10, 14, 15, 16, 18, 19, 20, 1, 5, 34, 72, 21} // shorthand declaration
	fmt.Println(arr3)
	fmt.Println("Length of an array:", len(arr3))
	fmt.Println("Capacity of an array:", cap(arr3))
	for _, val := range arr3 {
		print(val, " ")
	}

	err := changeVal(arr3, 3, 116)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arr3)
}

// nil can be used for pointers, slice, map, chan and interface,interface{}
func changeVal(arr [12]int, i uint, v int) error {
	if int(i) >= len(arr) {
		//return errors.New("invalid index")
		return fmt.Errorf("invald index")
	}
	arr[i] = v
	fmt.Println("Inside function", arr)
	return nil
}
