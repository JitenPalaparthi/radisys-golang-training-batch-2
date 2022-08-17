package main

import "fmt"

func main() {
	arr1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr1)

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			print(arr1[i][j], " ")
		}
		println()
	}

	arr2 := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			for k := 0; k < len(arr2[i][j]); k++ {
				print(arr2[i][j][k], " ")
			}
		}
		println()
	}
}

// Reverse an array
// find max and min elements from an array
// find the sum of elements in an array
