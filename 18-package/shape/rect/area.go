package rect

import "fmt"

func init() {
	fmt.Println("Init is called")
	Length = 1
	Width = 1
}

func init() {
	fmt.Println("Init is called 2nd time")
}

func init() {
	fmt.Println("Init is called 3rd time")
}

var (
	Length, Width float32
)

func AreaGlobal() float32 {
	return Length * Width
}
func Area(l, b float32) float32 {
	return l * b
}

func Greet() {
	fmt.Println("Hello World")
}
