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

	var a any = 100

	var b int = 200

	var c float32 = 300.3

	d := a.(int) + b + int(c)
	fmt.Println("a+b+c as int:", d)

	e := float32(a.(int)) + float32(b) + c

	fmt.Println("a+b+c as float32:", e)

	f := float64(float32(a.(int)) + float32(b) + c)
	//g := float64(a.(int)) + float64(b) + float64(c)
	fmt.Println("a+b+c as float64:", f)
}

// a as any  = 100
// b as int = 200
// c as float32 = 300.300

// a+b+c in int format
// a+b+c in float32 format
