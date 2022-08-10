package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	println("Hello World!")
	println("Hello", "Radisys minds!")
	fmt.Println("Hello", "World!")
	var str = "Hello World!" // Will discuss more in upcoming sections
	fmt.Printf("%s\n", str)
	fmt.Println("Current Date and Time:", time.Now())
	fmt.Println("Random number,:", rand.Intn(1999))
}
