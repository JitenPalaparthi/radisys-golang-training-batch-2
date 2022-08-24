package main

import "fmt"

func main() {

	var num1 myInt = 100

	fmt.Printf("Val of num1:%d Type of num1:%T\n", num1, num1)
	str1 := num1.ToString()
	fmt.Printf("Value of str1:%s Type of str1:%T\n", str1, str1)

	var num2 int = 200

	num1 = myInt(num2) // type casting is required
	fmt.Printf("Val of num1:%d Type of num1:%T\n", num1, num1)

	var num3 any = myInt(400)

	num1 = num3.(myInt)
	fmt.Printf("Val of num1:%d Type of num1:%T\n", num1, num1)

	var num4 int = int(num1)
	fmt.Printf("Val of num4:%d Type of num4:%T\n", num4, num4)

	num1.Increment()    //ok // based on the receiver it understand to perform call/pass by ref operation
	(&num1).Increment() //also ok
	fmt.Printf("Val of num1:%d Type of num1:%T\n", num1, num1)
	Increment(&num1) // have to pass pointer as a argument

}

// 1- create user defined type using struct
// 2- create user defined type from any existing type

type myInt int

func (mi *myInt) ToString() string {
	return fmt.Sprint(*mi)
}

func (mi *myInt) Increment() {
	*mi++
}

func Increment(mi *myInt) {
	*mi++
}

// create a new method calle ToBytes()[]byte
