package main

import "fmt"

// constats cannot be changed or reassigned
// constants must be evaluated at compile time.
// as long as the expression is constant (either direct value or from another constant it can be done)
func main() {
	const PI float32 = 3.14
	// PI = 3.45 // constants cannot be changed
	fmt.Println(PI)

	const (
		MIN, MAX = 1, 10
	)
	const SQUARE1 = 10 * 10 // compiler knows
	//var num = 10
	//const SQUARE2 = num * num // compiler does not evaluate becase num is a normal variable
	// and num can be changed because of any other expression/evaluation

	const SQUARE3 = MAX * MAX // compiler knows

	println(SQUARE1, SQUARE3)
}
