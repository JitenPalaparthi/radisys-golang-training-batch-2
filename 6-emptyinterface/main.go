package main

import (
	"fmt"
	"reflect"
)

// type any = interface{}
// cannot do type casting from empty interface to a type
// have to do type assertion
// type caste:     T(v)
// type assertion: v.(T)
func main() {

	//var val interface{}
	var val any
	// Object obj

	val = 100
	fmt.Println("Value of val:", val, "Type of val:", reflect.TypeOf(val))
	var num1 int = val.(int) // type assertion
	fmt.Println("Value of num1:", num1, "Type of num1:", reflect.TypeOf(num1))

	val = false
	fmt.Println("Value of val:", val, "Type of val:", reflect.TypeOf(val))
	boolean := val.(bool)
	fmt.Println("Value of boolean:", boolean, "Type of boolean:", reflect.TypeOf(boolean))

	val = "Hello World"
	fmt.Println("Value of val:", val, "Type of val:", reflect.TypeOf(val))

	val = 12.45
	fmt.Println("Value of val:", val, "Type of val:", reflect.TypeOf(val))

	val = 'A'
	fmt.Println("Value of val:", val, "Type of val:", reflect.TypeOf(val))

	val = byte(95)
	fmt.Println("Value of val:", val, "Type of val:", reflect.TypeOf(val))

}

// a as any  = 100
// b as int = 200
// c as float32 = 300.300

// a+b+c in int format
// a+b+c in float32 format
