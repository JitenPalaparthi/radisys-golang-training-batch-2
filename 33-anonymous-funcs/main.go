package main

import "fmt"

type calcFunc func(int, int) int

func main() {

	fmt.Println("store and call func from a map")

	var funcMap map[string]calcFunc

	funcMap = make(map[string]calcFunc)

	addition := func(a, b int) int {
		return a + b
	}

	funcMap["add"] = addition

	funcMap["sub"] = sub

	funcMap["mul"] = func(a, b int) int {
		return a * b
	}

	funcMap["div"] = func(a, b int) int {
		return a / b
	}

	a, b := 20, 10
	for key, f := range funcMap {
		fmt.Println(key, ":", f(a, b))
	}
	fmt.Println("\nstore and call func from a slice")
	slice := make([]calcFunc, 4)
	//var slice [4]calcFunc
	slice[0] = addition
	slice[1] = sub
	slice[2] = func(a, b int) int {
		return a * b
	}
	slice[3] = func(a, b int) int {
		return a / b
	}

	//a, b := 20, 10 already declared above so not required
	for i, f := range slice {
		fmt.Println(i, ":", f(a, b))
	}

}

func sub(a, b int) int {
	return a - b
}
