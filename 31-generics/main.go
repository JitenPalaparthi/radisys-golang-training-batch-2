package main

import "fmt"

func main() {

	c1 := &Calc[int]{A: 10, B: 20}
	fmt.Println("c1 add:", c1.Add())
	fmt.Println("c1 sub:", c1.Sub())
	fmt.Println("c1 mul:", c1.Mul())
	fmt.Println("c1 div:", c1.Div())
	println()

	c2 := &Calc[float64]{A: 10.1, B: 20.2}
	fmt.Println("c2 add:", c2.Add())
	fmt.Println("c2 sub:", c2.Sub())
	fmt.Println("c2 mul:", c2.Mul())
	fmt.Println("c2 div:", c2.Div())
}

type Type interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Calc[T Type] struct {
	A, B T
}

func (c *Calc[T]) Add() T {
	return c.A + c.B
}

func (c *Calc[T]) Sub() T {
	return c.A - c.B
}

func (c *Calc[T]) Mul() T {
	return c.A * c.B
}

func (c *Calc[T]) Div() T {
	return c.A / c.B
}

// type SType interface {
// 	string | []byte
// }
// type Calc[T Type, S SType] struct {
// 	A, B T
// 	C, D S
// }

// func (c *Calc[T, S]) Add() T {
// 	return c.A + c.B
// }

// func (c *Calc[T, S]) Sub() T {
// 	return c.A - c.B
// }

// func (c *Calc[T, S]) Mul() T {
// 	return c.A * c.B
// }

// func (c *Calc[T, S]) Div() T {
// 	return c.A / c.B
// }
