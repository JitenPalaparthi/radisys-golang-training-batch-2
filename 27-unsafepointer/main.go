package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var arr1 [5]int = [5]int{11, 12, 13, 14, 15}

	ptr1 := &arr1[0]

	size1 := unsafe.Sizeof(arr1[0])

	var arr2 [4]string = [4]string{"hi", "doing.How is your work.Are you able to learn new stuff", "are", "you"}
	ptr2 := &arr2[0]
	size2 := unsafe.Sizeof(arr2[0])
	fmt.Println("Size of string:", size2)
	//ptr1 += size // go does not suppport pointer arthemetic

	// arrptr := uintptr(unsafe.Pointer(ptr1))

	// // fmt.Println(arrptr, "-->", size)
	// fmt.Printf("0x%x\n", arrptr)
	// // fmt.Println("\n", ptr1)

	// arrptr += size // arrptr = arrptr + 8 (assume on 64 bit arch type)
	// fmt.Printf("0x%x\n", arrptr)

	// val := (*int)(unsafe.Pointer(arrptr))
	// fmt.Println("2nd Element in an array:", *val)
	arrptr1 := uintptr(unsafe.Pointer(ptr1))
	fmt.Println(*ptr1)
	for i := 1; i < 15; i++ {
		arrptr1 += size1
		val := (*int)(unsafe.Pointer(arrptr1))
		fmt.Println(*val)
		*val = *val + 100

	}

	arrptr2 := uintptr(unsafe.Pointer(ptr2))
	fmt.Println(*ptr2)
	for i := 1; i < 4; i++ {
		arrptr2 += size2
		val := (*string)(unsafe.Pointer(arrptr2))
		fmt.Println(*val)
		*val = *val + " Can I dereference you?"
	}

	// arrptr1 = uintptr(unsafe.Pointer(ptr1))
	// fmt.Println(*ptr1)
	// for i := 1; i < 4; i++ {
	// 	arrptr1 += size1
	// 	val := (*int)(unsafe.Pointer(arrptr1))
	// 	fmt.Println(*val)
	// }

	//fmt.Println(arr1)
	fmt.Println(arr2)
}

// builtin type in golang
// builtin type package in golang

// CV library written in c lang want to use it in go
// Kafka go client --> c library imported to go
// ml algorithms.. c library imported to go
// import python library in go --> export from python to c import in go
// python -> c -> go
// java -> c -> go

// *void
// return pointer instead of arr
// pass pointers in c

// getKeys()[]string
// getKeys() *void
