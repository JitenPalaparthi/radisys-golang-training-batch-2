package main

import "fmt"

func main() {

	fmt.Println("sumOfInts:", sumOfInts(1, 2, 3, 4, 5))
	fmt.Println("sumOfAny:", sumOfAny(1, 2, 3, 4, 5))
	fmt.Println("sumOfAny:", sumOfAny(1, 2, 3, 4, 5.0, true))
	fmt.Println("sum:", sum(1, 2, 3, 4, 5))
	fmt.Println("sum:", sum(1.1, 2.1, 3.1, 4.1, 5.1))

	//fmt.Println("sum:", sum(1.1, 2.1, 3.1, 4.1, 5.1, true))

	var a, b, c, d, e, f float = 1.1, 2.1, 3.1, 4.1, 5.1, 6.1
	s := sumUType(a, b, c, d, e, f)
	fmt.Println("Sum of floats in string:", s.ToString())

}

func sumOfInts(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
func sumOfAny(vals ...any) float64 {
	sum := 0.0
	for _, v := range vals {
		switch v := v.(type) {
		case int:
			sum += float64(v)
		case uint8:
			sum += float64(v)
		case uint16:
			sum += float64(v)
		case uint32:
			sum += float64(v)
		case uint64:
			sum += float64(v)
		case int8:
			sum += float64(v)
		case int16:
			sum += float64(v)
		case int32:
			sum += float64(v)
		case int64:
			sum += float64(v)
		case float32:
			sum += float64(v)
		case float64:
			sum += v
		}

	}
	return sum
}

func sumM[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](nums ...T) T {
	var s T
	for _, v := range nums {
		s += v
	}
	return s
}

func sum[T Type](nums ...T) T {
	var s T
	for _, v := range nums {
		s += v
	}
	return s
}

func sumUType[T UType](vals ...T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}

type Type interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type float float64

func (f float) ToString() string {
	return fmt.Sprint(f)
}

type UType interface {
	float
	ToString() string
}

// type Calc struct
// A,B fields must be generic types
// Add,Sub,Mul,Div are methods for Calc
