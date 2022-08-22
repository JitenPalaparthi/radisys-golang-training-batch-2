package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a, b = 10, 20
	println("a:", a, "b:", b)
	c := add(a, b)
	println("c:", c)

	var a1, b1, c1 any = 10, 20, 0
	c1 = add(a1.(int), b1.(int))
	println("c1:", c1)

	slice1 := []int{10, 20, 30, 40}
	println(slice1)
	//---------------------------------
	c2 := sum(slice1)
	println("c2:", c2)
	// dangling pointer inside sumOf
	c3 := sumOf(slice1)
	println("c3:", *c3)

	slice2 := make([]int, 100)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = rand.Intn(20000)
	}
	c4 := sumOf(slice2)
	println("c3:", *c4)

	str1 := "Hello World!"
	str1 = str1 + ". How are you doing"
	println(str1)

	fmt.Println("Address of c3 and c4", c3, c4)

	c5 := new(int)
	*c5 = 500
	println("C5:", *c5)

}

func add(a, b int) int {
	return a + b
}

func sum(slice []int) int {
	s := 0
	for _, v := range slice {
		s += v
	}
	return s
}

func sumOf(slice []int) *int {
	s := 0
	for _, v := range slice {
		s += v
	}
	return &s // Dangling pointer in c lang
}

// does not escape --> to store in stack
// escape to heap  --> compiler has decided to create this varailbe in heap
// moved to heap   --> runtime has decided to move certain data/variable to heap
