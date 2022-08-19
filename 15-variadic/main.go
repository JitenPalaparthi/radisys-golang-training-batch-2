package main

import "fmt"

func main() {

	fmt.Println("Hello", "World", ".", "How", "are", "you", "doing", "?")
	fmt.Println("Hello World. How are you doing ?")
	// variadic parameter

	fmt.Println("Sum:", sumOf(10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))
	fmt.Println("Sum:", sumOf(10, 11, 12))
	fmt.Println("Sum:", sumOf())

	slice1 := []int{}
	slice1 = append(slice1, 10, 20, 30, 40, 50)

	fmt.Println("Sum:", sumOf(slice1...))

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Sum:", sumOf(arr1[:]...))

	i := 2
	v := 120
	// [10, 20, 30][30, 40, 50]
	// [10 20 30 30 40 50]
	slice1 = append(slice1[0:i], slice1[i-1:]...)
	slice1[i] = v
	//[10 20 120 30 40 50]
	fmt.Println("slice1", slice1)

}

// ...Type to be given as a variadic paramter
// variadic parameter must be the last parameter
func sumOf(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// sumOfAny(10,10.5,200,200.45,10,11,12,34.5,.5,.6,.7)
func sumOfAny(vals ...any) (any, error) {

	return nil, nil
}

// delete an element from the slice
// task-1
func delete(arr []int, i int) []int {

	return nil
}
