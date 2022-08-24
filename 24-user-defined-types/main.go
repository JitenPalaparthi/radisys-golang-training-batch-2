package main

import (
	"fmt"
	"shapes/shape/cube"
	"shapes/shape/rect"
	"shapes/shape/square"
)

func main() {

	p1 := Person{ID: 101, Name: "Jiten", Email: "JitenP@Outlook.Com"}
	Display(p1)
	p1.Display() // calling method of p1

	// Shape calls

	//c1 := &cube.Cube{L: 12.45, B: 13.45, H: 15.67}
	c1, err := cube.New(12, 45, 13.46, 15.67)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Area of Cube:", c1.Area())
	fmt.Println("Perimeter of Cube:", c1.Perimeter())

	fmt.Println("Area of Cube stored:", c1.A)
	fmt.Println("Perimeter of Cube stored:", c1.P)

	r1 := &rect.Rect{L: 12.45, B: 13.46}
	fmt.Println("Area of Rect:", r1.Area())
	fmt.Println("Perimeter Of Rect:", r1.Perimeter())

	s1 := &square.Square{S: 25.25}

	fmt.Println("Area of Square:", s1.Area())
	fmt.Println("Perimeter of Square:", s1.Perimeter())
}

// shapes
// 		shape
//			cube
//				cube.go
//			rect
//				rect.go
//			square
//				square.go

type Person struct {
	ID          int
	Name, Email string
}

// function
func Display(p Person) {
	fmt.Println("ID:", p.ID)
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
}

// Method of Person object
// method contains receiver --> (p Person)
func (p Person) Display() {
	fmt.Println("ID:", p.ID)
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
}

type T1 struct{}

func (t1 T1) Display() {
	fmt.Println("This is an empty struct")
}

// Calc
// New function
// Add Method
// Sub Method
// Multiply Method
// Divide Method

// any type. all must be like methods not like functions
