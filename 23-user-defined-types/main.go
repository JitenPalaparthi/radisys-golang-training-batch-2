package main

import "fmt"

func main() {
	var c1 ColourCode
	c1.int = 101
	c1.string = "Red"
	c1.bool = true
	fmt.Println(c1)

	var e1 EmptyStruct
	fmt.Println(e1)
}

type ColourCode struct {
	int
	string
	bool
}

type EmptyStruct struct{}

// create Employee struct
// Address struct
