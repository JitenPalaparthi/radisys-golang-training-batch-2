package main

import (
	"fmt"
	"reflect"
)

func main() {
	c1 := complex(12.14, 13.45)

	var (
		r1, i1         = 12.14, 13.45
		r2, i2 float32 = 12.14, 13.45
	)
	c2 := complex(r1, i1)
	c3 := complex(r2, i2)
	c4 := complex(r2, 13.45)
	c5 := 12.14 + 13.45i // operator overloading of complex numbers

	fmt.Println("Value of c1:", c1, "Type of c1:", reflect.TypeOf(c1))
	fmt.Println("Value of c2:", c2, "Type of c2:", reflect.TypeOf(c2))
	fmt.Println("Value of c3:", c3, "Type of c3:", reflect.TypeOf(c3))
	fmt.Println("Value of c4:", c4, "Type of c4:", reflect.TypeOf(c4))
	fmt.Println("Value of c5:", c5, "Type of c5:", reflect.TypeOf(c5))
	// bigint there is a special package for that purpose

	fmt.Println("Addition of complex numbers:", c1+c2)
	fmt.Println("Subscraction of complex numbers:", c1-c2)
	fmt.Println("Multiplication of complex numbers:", c1*c2)
	fmt.Println("Division of complex numbers:", c1/c2)

}
