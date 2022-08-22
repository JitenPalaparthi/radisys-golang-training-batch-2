package main

import "fmt"

// How to create a pointer
// nil pointer
// dereference a pointer
// passing pointers as arguments to a function call
// return pointer
// creating pointer to a pointer
// creating a pointer using new function
// type infernece in pointer

// & address of
// * value at address

// 1. Pointer is to hold address or another variable
// 2. pointer is to hold address

func main() {
	var num1 int = 100
	var ptr1 *int = &num1 // ptr1 does not contain any address or address of any varaible hence it is nil
	if ptr1 == nil {
		println("its  nil pointer")
	}

	//ptr1 = &num1 // ptr1 is not nil, it is having an address. Address of num1

	fmt.Println("Value of num1", num1, "address of num1", &num1)
	fmt.Println("Value of num1 through ptr1", *ptr1, "address of num1 that ptr1 holds", ptr1)

	*ptr1 = 101 // derefernece

	fmt.Println("Value of num1", num1, "address of num1", &num1)
	fmt.Println("Value of num1 through ptr1", *ptr1, "address of num1 that ptr1 holds", ptr1)

	fmt.Println("Address of ptr1:", &ptr1, "address that pointer hosts", ptr1, "value that pointer holds", *ptr1)

	var ptr2 **int

	ptr2 = &ptr1

	fmt.Println("Value of num1 through ptr2", **ptr2)

	**ptr2 = 102

	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr1", *ptr1)
	fmt.Println("Value of num1 through ptr2", **ptr2)

	changeVal(&num1, 103)
	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr1", *ptr1)
	fmt.Println("Value of num1 through ptr2", **ptr2)

	changeVal(ptr1, 104)
	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr1", *ptr1)
	fmt.Println("Value of num1 through ptr2", **ptr2)

	changeVal(*ptr2, 105)
	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr1", *ptr1)
	fmt.Println("Value of num1 through ptr2", **ptr2)

	var ptr3 ***int
	ptr3 = &ptr2
	fmt.Println("Value of num1", ***ptr3)

	changeVal(**ptr3, 106)
	fmt.Println("Value of num1", num1)
	fmt.Println("Value of num1 through ptr1", *ptr1)
	fmt.Println("Value of num1 through ptr2", **ptr2)
	fmt.Println("Value of num1 through ptr3", ***ptr3)
	// when will can you say a pointer is nil
	// if a pointer does not hold any memory then we can call it as nil pointer
	ptr4 := new(int) // this is not a nil pointer. a pointer is instantiated using new built in function
	// The above code holds memory but not through a valriable
	println(ptr4)
	println(*ptr4) // type inference works
	*ptr4 = 100
	println(*ptr4) // type inference works

	var num2 int = 200

	ptr5 := &num2 // short hand declaration of pointer

	println("Pointer declared shorthand notation", *ptr5)

	// ptr5++ // cant do pointer arthimetic directly
}

func changeVal(num *int, val int) {
	*num = val
}

// 1- create a calc pacakge
// 2- create add function add(*int,*int)(*int)
// 3- create sub fucntion sub(*int,*int)(*int)
// create those two functions with pointers as parameters.
