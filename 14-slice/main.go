package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	/*
		    Slice
			---------
			Ptr:        0x00543fd
			Length:     10
			Capacity:   20
			----------

			When you call append
	*/

	fmt.Println("Len:", len(slice1), "Cap:", cap(slice1), "slice1:", slice1, "addressOf", &slice1[0]) // len: 10 and cap: 10
	slice1 = insert(slice1, 8, 88)
	fmt.Println("Len:", len(slice1), "Cap:", cap(slice1), "slice1:", slice1)

	slice2 := make([]int, 10, 20)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = i + 1
	}
	fmt.Println("Len:", len(slice2), "Cap:", cap(slice2), "slice2:", slice2, "addressOf", &slice2[0])
	slice2 = insert(slice2, 8, 88)
	fmt.Println("Len:", len(slice2), "Cap:", cap(slice2), "slice2:", slice2)

	// usually all assigns are reference copy of a slice.

	// deep copy

	slice3 := make([]int, len(slice2)) // deep copy both slices are two different memory locations.. two different copies
	// for i, v := range slice2 {
	// 	slice3[i] = v
	// }

	copy(slice3, slice2)
	fmt.Println("Len:", len(slice3), "Cap:", cap(slice3), "slice3:", slice3)

}

// insert operation in slices
func insert(arr []int, i, v int) []int {
	if i != 0 {
		arr = append(arr[:i], arr[i-1:]...)
		arr[i] = v
		return arr
	} else {
		arr = append(arr[:0], arr[i:]...)
		arr[i] = v
		return arr
	}

	//fmt.Println(&arr[0], arr)
}

// delete an element from the slice
// task-1
func delete(arr []int, i int) []int {

	return nil
}
