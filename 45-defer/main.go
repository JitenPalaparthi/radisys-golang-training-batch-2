package main

import "fmt"

func main() {
	//1
	{
		defer fmt.Println("End of main")
	}
	//2
	defer println("call-1")
	//3
	defer println("call-2")
	//4
	defer println("call-3")
	fmt.Println("Main started")
}

// 4
// 3
// 2
// 1
