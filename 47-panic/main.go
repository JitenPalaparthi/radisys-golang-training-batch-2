package main

import (
	"fmt"
)

// panic panics the current execution
// defers are called when func panics

func main() {
	// i := 0
	// fmt.Println(100 / i)
	defer fmt.Println("end of main")
	fn := new(string)
	*fn = "Jiten"
	ln := new(string)
	*ln = "p"
	finalName(fn, ln)
	finalName(nil, nil)
	greet()
	//runtime.Goexit()
	// end of main func to be executed here and hence defer
}

func greet() {
	fmt.Println("Hello World")
}

func finalName(fn *string, ln *string) {
	if fn == nil {
		panic("firstName is nil")
	}
	if ln == nil {
		panic("lastName is nil")
	}
	fmt.Println(*fn, *ln)
}
