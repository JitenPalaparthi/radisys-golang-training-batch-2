package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var age uint8 = 28

	if age >= 18 {
		println("age is", age, ".So eligible for vote")
	} else {
		println("age is", age, ".So not eligible for vote")
	}
	gender := 'M'
	// if age >= 21 {
	// 	println("Age is", age, "gender is", string(gender), ".Eligible for marriage")
	// } else
	if (age >= 18 && gender == 'F') || (age >= 21 && gender == 'M') {
		println("Age is", age, "gender is", string(gender), ".Eligible for marriage")
	} else {
		println("not eligible for marraige")
	}

	if num := rand.Intn(2001); num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
		//return
	}

	ok := true
	if !ok {
		fmt.Println("true")
	} else {
		fmt.Print("false")
	}
}
