package main

import "fmt"

func main() {
	// func --> to create a func
	// func() --> input parameters
	// { } --> body of the function
	// () --> to execute the func

	// func no parameters, no return types
	func() {
		fmt.Println("Hello World!-1")
	}()

	// func with string as input parameter ;no return types
	func(msg string) {
		fmt.Println(msg)
	}("Hello World!-2")

	// func with string as input parameter and string as return type

	str := func(msg string) string {
		return msg
	}("Hello World!-3")

	fmt.Println(str)

	// sum of two numbers: two input parameters and one output parameter

	sum := func(a, b int) int {
		return a + b
	}
	c := sum(10, 20) // execution of func
	fmt.Println("Sum:", c)

	fmt.Println("\n calc calls")
	c1 := calc(10, 20, sum)
	fmt.Println("Sum:", c1)

	c2 := calc(20, 10, func(a, b int) int {
		return a - b
	})
	fmt.Println("Sub:", c2)

	c3 := calc(20, 10, func(a, b int) int {
		return a * b
	})
	fmt.Println("Mul:", c3)

	c4 := calc(20, 10, func(a, b int) int {
		return a / b
	})
	fmt.Println("Div:", c4)

	// calc1

	fmt.Println("\ncalc1 calls")
	c5 := calc1(10, 20, sum)
	fmt.Println("Sum:", c5)

	c6 := calc1(20, 10, func(a, b int) int {
		return a - b
	})
	fmt.Println("Sub:", c6)

	c7 := calc1(20, 10, func(a, b int) int {
		return a * b
	})
	fmt.Println("Mul:", c7)

	c8 := calc1(20, 10, func(a, b int) int {
		return a / b
	})
	fmt.Println("Div:", c8)

}

func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}

type calcFunc func(int, int) int

func calc1(a, b int, cf calcFunc) int {
	return cf(a, b)
}
