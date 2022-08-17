package main

import (
	"fmt"
	"math/rand"
)

// slice is similar to array but can be expanded at runtime
// slice,map, chan --> make function
func main() {
	var slice1 []int // only declaration
	// type inference has given [0,0,0,0,0] as default values
	slice1 = make([]int, 5) // instantiate []int with the length of 5 and capacity of 5 elements
	fmt.Println(slice1)

	for i := range slice1 {
		slice1[i] = rand.Intn(1000)
	}
	fmt.Println(slice1, "length:", len(slice1), "capacity", cap(slice1))

	slice2 := []int{112, 234, 89, 231, 323, 43, 3545, 6, 43, 45, 6566, 54, 3} // shorhad declaration and value assignment
	fmt.Println(slice2, "length:", len(slice2), "capacity", cap(slice2))

	slice3 := make([]int, 1) // [0,0,0,0,0]
	fmt.Println(slice3, "length:", len(slice3), "capacity", cap(slice3), &slice3[0])
	slice3 = append(slice3, 100)
	fmt.Println(slice3, "length:", len(slice3), "capacity", cap(slice3), &slice3[0])
	changeVal(slice2, 1, 200) // slice is a reference type
	fmt.Println(slice2, "length:", len(slice2), "capacity", cap(slice2))
	slice4 := slice2 // reference copy not a deep copy

	slice4[2] = 300 // since slice4 is refefence of slice2 .. change in slice4 changes slice2 as well
	// becasue both are pointed to same memory
	fmt.Println(slice2, "length:", len(slice2), "capacity", cap(slice2))

	min, max, err := getMinAndMax(slice2)
	fmt.Println("Min:", min, "Max:", max, "-->", slice2)
	var slice5 []int // declared but not instantiated
	min, max, err = getMinAndMax(slice5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Min:", min, "Max:", max, "-->", slice5)
	}

	arr1 := [...]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("arr1:", arr1, "length:", len(arr1), "capacity:", cap(arr1))
	changeVal(arr1[:], 1, 200) // convert array to slice
	fmt.Println("arr1:", arr1, "length:", len(arr1), "capacity:", cap(arr1))

	arr2 := arr1[:]
	arr2 = append(arr2, 21)
	fmt.Println("arr2:", arr2, "length:", len(arr2), "capacity:", cap(arr2))

	fmt.Println("All elements", arr2[:])
	fmt.Println("0th to the 5th elements", arr2[:5])
	fmt.Println("5th to the end elements", arr2[5:])
	fmt.Println("5th to the 8th elements", arr2[5:8])

	//[10 200 112]  [12 13 14 15 16 17 18 19 20 2]
	arr2 = append(arr2[:3], arr2[2:]...)
	arr2[2] = 122

	fmt.Println(arr2)
	insert(arr2, 0, 4000)
	fmt.Println(arr2)

}

// insert operation in slices
func insert(arr []int, i, v int) {
	if i != 0 {
		arr = append(arr[:i], arr[i-1:]...)
		arr[i] = v
	} else {
		arr = append(arr[:0], arr[i:]...)
		arr[i] = v
	}
}

func changeVal(arr []int, i uint, v int) error {
	if int(i) >= len(arr) {
		return fmt.Errorf("invalid index")
	}
	arr[i] = v
	return nil
}

func getMinAndMax(arr []int) (min int, max int, err error) {
	if arr == nil {
		return min, max, fmt.Errorf("nil array")
	}

	if len(arr) == 1 {
		return arr[0], arr[0], nil
	}

	min = arr[0]
	max = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, max, nil
}

// palendrome
// transpose matrix
// delete an element from a slice
