package main

import "fmt"

var i = 100

func main() {
	num := new(int)
	*num = 100
	//i := 100
	i = i + 100
	defer fmt.Println("Calling in defer", i)

	*num = *num + 100
	defer fmt.Println("Calling ptr in defer", *num)
	// defer func(i int) {
	// 	i = i + 100
	// 	fmt.Println("called widh defer", i)
	// }(i)

	i = i / 100
	fmt.Println("normal call", i)

	*num = *num / 100
	fmt.Println("normal ptr call", *num)

	// defer is called here during runtime

	str := "Hello World"

	for _, v := range str {
		defer fmt.Print(string(v))
	}
}
