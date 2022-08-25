package main

import (
	"fmt"
	"shapes/shape"
	"shapes/shape/rect"
	"shapes/shape/square"
)

func main() {

	// var iShape shape.IShape
	// iShape = &rect.Rect{L: 22.45, B: 28.95}
	// Area(iShape)
	// Perimeter((iShape))
	// println()

	// var s1 square.Square = 25.25
	// iShape = &s1

	// Area(iShape)
	// Perimeter(iShape)

	r1 := &rect.Rect{L: 10.57, B: 12.93}
	if shape1, err := shape.New(r1); err != nil {
		fmt.Println(err)
	} else {
		shape1.Area()
		shape1.Perimeter()
	}

	s1 := new(square.Square)
	*s1 = 25.25

	if shape2, err := shape.New(s1); err != nil {
		fmt.Println(err)
	} else {
		shape2.Area()
		shape2.Perimeter()
	}

}

// func Area(iShape shape.IShape) {
// 	fmt.Println("Area:", iShape.Area())
// }
// func Perimeter(iShape shape.IShape) {
// 	fmt.Println("Perimeter:", iShape.Perimeter())
// }

// add cuboid, triangle to this
// create a slice of  []Shape
// create instances of rect,Square,Cuboid, Triangle
// pass them as parameters for Shape
// use range loop to call Area and Perimeter methods of each Shape object
