package main

import (
	"fmt"
	"reflect"
)

// variable delcartion
func main() {
	var num1 int
	// type inference
	// based on the type a value is given to the variable
	// like clang go does not take garbage values / random values
	fmt.Println("Value of num1:", num1, " type of num1:", reflect.TypeOf(num1))
	var ok bool
	fmt.Printf("Value of ok: %v  type of ok: %T", ok, ok)

	var num2 float32
	fmt.Println("\nValue of num2:", num2, " type of num2:", reflect.TypeOf(num2))

	var num3 float64 = 1000.100001
	fmt.Println("Value of num3:", num3, " type of num3:", reflect.TypeOf(num3))

	var bt byte = 95 // assigning value
	fmt.Println("\nValue of bt:", bt, " type of bt", reflect.TypeOf(bt))
	fmt.Printf("Value of bt: %v type of bt:%T", bt, bt)

	var rn rune = '喂' //19000 // The size of the rune is 32bit
	fmt.Println("\nValue of rn:", rn, " type of rn", reflect.TypeOf(rn))

	var str string = "Hello Wrold!"
	fmt.Println("Value of str:", str, " type of str", reflect.TypeOf(str))

	var num4 = 100 // int

	var num5 = 100.101 // float64

	var done = false // bool

	var name = "Jiten" //str

	fmt.Println("Value of num4:", num4, " type of num4", reflect.TypeOf(num4))
	fmt.Println("Value of num5:", num5, " type of num5", reflect.TypeOf(num5))
	fmt.Println("Value of done:", done, " type of done", reflect.TypeOf(done))
	fmt.Println("Value of name:", name, " type of name", reflect.TypeOf(name))

	var ( // no type inference. Value to be given. Then only compiler can evaluate and give a type at comptime
		a, b, c, d = 10, 20.20, false, "Hello"
	)

	fmt.Println(a, b, c, d)

}
