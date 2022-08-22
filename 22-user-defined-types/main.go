package main

import "fmt"

func main() {

	// anonymous struct

	p1 := struct {
		ID    int
		Name  string
		Email string
	}{
		ID:    101,
		Name:  "John",
		Email: "JitenP@Outlook.Com",
	}

	fmt.Println(p1)

	p2 := struct {
		ID    int
		Name  string
		Email string
	}{
		ID:    102,
		Name:  "John",
		Email: "John.Matt@Outlook.Com",
	}
	fmt.Println(p2)

}
