package main

import (
	"fmt"
	"unsafe"
)

func main() {
	t1 := T1{id: 101, name: "jiten", isMarried: true, email: "jitenp@outlook.com", age: 38}
	t2 := T2{id: 101, name: "jiten", isMarried: true, email: "jitenp@outlook.com", age: 38}
	fmt.Println(t1, "Sizeof:", unsafe.Sizeof(t1))
	fmt.Println(t2, "Sizeof:", unsafe.Sizeof(t2))
	t3 := T3{ok: true}
	fmt.Println(t3, "Sizeof:", unsafe.Sizeof(t3))

}

type T1 struct {
	id        uint16 // 8   --> 2	 6
	name      string // 16  --> 0    0
	ref       uint32 // 8   --> 4    4
	addr      string // 16
	isMarried bool   // 8   --> 1    7
	email     string // 16  --> 0    0
	age       uint8  // 8   --> 1    7
} //16 bytes

type T2 struct {
	id        uint16 // 8 --> 2 	6
	age       uint8  // 6 --> 1     5
	isMarried bool   // 5 --> 1     4
	ref       uint32
	name      string // 16
	email     string // 16
	addr      string
}

type T3 struct {
	no      uint16
	ok      bool
	done    bool
	ref     uint32
	name    string
	another bool
}
