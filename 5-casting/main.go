package main

import (
	"fmt"
	"reflect"
	"strconv"
	_ "time" // blank identifier
)

// type casting
// there is not implicit caste in golang
// no type automatically converted to another type
func main() {
	var age1 uint8 = 39
	var age2 uint16 = uint16(age1) // not XXX (uint16)age1

	fmt.Println("Value of age1:", age1, "Type of age1:", reflect.TypeOf(age1))
	fmt.Println("Value of age2:", age2, "Type of age2:", reflect.TypeOf(age2))

	var char rune = '䨸' // Internally rune is nothing but int32
	var charVal int32 = char
	fmt.Println("Value of charVal:", charVal, "Type of charVal:", reflect.TypeOf(charVal))

	var byt uint8 = 'A'
	fmt.Println("Value of byt:", byt, "Type of byt:", reflect.TypeOf(byt))
	fmt.Println(string(19000))

	var val float32 = 12.34

	var num int = int(val)
	fmt.Println("Value of num:", num, "Type of num:", reflect.TypeOf(num))
	var long int64 = 456067974523232

	var tiny uint8 = uint8(long)
	fmt.Println("Value of tiny:", tiny, "Type of tiny:", reflect.TypeOf(tiny))

	// type conversion

	num1 := 65

	str1 := string(num1) // not "65"
	println(str1)

	//str2 := fmt.Sprintf("%v", num1)
	str2 := fmt.Sprint(num1)
	println(str2)

	str3 := strconv.Itoa(num1)
	println(str3)

	str4 := "19000"
	num2, _ := strconv.Atoi(str4) // _ blank identifier
	println(num2)

	//  swap

	a, b, c := 100, 200, 300
	fmt.Println("a:", a, "b:", b, "c:", c)
	// t := a
	// a = b
	// b = t
	a, b, c = c, a, b // easy swapping
	fmt.Println("a:", a, "b:", b, "c:", c)

}
