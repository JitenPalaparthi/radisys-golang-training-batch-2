package main

import "fmt"

func main() {
	day := 15
	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	num := 500
	switch {
	case num < 50:
		println(num, "is less than 50")
	case num >= 50 && num < 100:
		println(num, "is greater than or equal to 50 and less than 100")
	case num >= 100 && num < 200:
		println(num, "is greater than or equal 100 and less than 200")
	default:
		println("num is greater than or equal 200")
	}

	str := "Hello World!"
	count := 0
	for _, char := range str {
		switch char {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			count++
			println(string(char), "is oval")
		default:
			println(string(char), "is either consolent or a special char")
		}
	}

	println("in", str, "there are", count, " ovals")

	num = 64

	// if you remove break in other programming languages same thing if keep fallthrough in go
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	var iface interface{}
	var ok bool = false
	iface = ok

	switch iface.(type) {
	case int:
		fmt.Println(iface, "is int type")
	case int8:
		fmt.Println(iface, "is int8 type")
	case int16:
		fmt.Println(iface, "is int16 type")
	case int32:
		fmt.Println(iface, "is int32 type")
	case int64:
		fmt.Println(iface, "is int64 type")
	case uint8:
		fmt.Println(iface, "is uint8 type")
	case uint16:
		fmt.Println(iface, "is uint16 type")
	case uint32:
		fmt.Println(iface, "is uint32 type")
	case uint64:
		fmt.Println(iface, "is uint64 type")
	case float32:
		fmt.Println(iface, "is float32 type")
	case float64:
		fmt.Println(iface, "is float64 type")
	case bool:
		fmt.Println(iface, "is boolean type")
	}
}
