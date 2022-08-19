package main

import (
	"fmt"
	"math/rand"
	r "shapes/shape/rect"
	"shapes/shape/square"
)

func main() {
	r.Greet()
	fmt.Println("Rand:", rand.Intn(1000))
	a1 := r.Area(10.4, 20.45)
	fmt.Println("Area of rect:", a1)
	p1 := r.Perimeter(10.4, 20.45)
	fmt.Println("Perimeter of Rect:", p1)

	a2 := square.Area(10.4)
	fmt.Println("Area of Square:", a2)
	p2 := square.Perimeter(10.4)
	fmt.Println("Perimeter of Square:", p2)

	// r.Length = 10.45
	// r.Width = 12.65

	a3 := r.AreaGlobal()
	p3 := r.PerimeterGlobal()

	fmt.Println("Area of Rect", a3)
	fmt.Println("Perimeter of Rect", p3)

}

// go does not have any access specifiers/modifiers
// any member of a package can be exported or not exported
// anything that starts with uppercase can be exported
// anything that statrs with lowercase cannot be exported
