package main

import "fmt"

func main() {
	var p1 Person
	p1.ID = 101
	p1.Name = "Jiten"
	p1.Email = "JitenP@Outlook.Com"
	p1.Contact = "9618558500"
	p1.SocialMedia = []string{"Jiten.Pala@Facebook", "Jiten_1984", "linkedin.com/jpalaparthi"}

	// p1.Address.Line1 = "1st Lane"
	// p1.Address.Street = "Kesavadasapuram"
	// p1.Address.City = "Trivandrum"
	// p1.Address.State = "Kerala"
	// p1.Address.Country = "India"
	// p1.Address.PIN = "590064"

	p1.A.Line1 = "1st Lane"
	p1.A.Street = "Kesavadasapuram"
	p1.A.City = "Trivandrum"
	p1.A.State = "Kerala"
	p1.A.Country = "India"
	p1.A.PIN = "590064"

	fmt.Println(p1)

	p2 := Person{ID: 102, Name: "Rahim", Email: "RahimSheik@Gmail.Com", Contact: "9618558500", SocialMedia: []string{"Rahim_Shiek"}}
	fmt.Println(p2)

	e1 := Employee{} // shorthand declaration
	e1.ID = 1007
	e1.Name = "Jiten"
	e1.Email = "JitenP@Outlook.Com"
	e1.Contact = "9618558500"
	// can access promoted fields through e1
	e1.Address.ID = 1 // becase Employee and Address both has ID field
	e1.Line1 = "1st Lane"
	e1.Street = "Kesavadasapuram"
	e1.City = "Trivandrum"
	e1.State = "Kerala"
	e1.Country = "India"
	e1.PIN = "590064"

	fmt.Println(e1)
}

type Person struct {
	ID                   int
	Name, Email, Contact string
	SocialMedia          []string
	A                    Address // composition
	//Address              Address // composition
}

type Address struct {
	ID                                       int
	Line1, Street, City, State, Country, PIN string
}

// Can assume Employee is dereived from Address
type Employee struct {
	ID                   int
	Name, Email, Contact string
	Address              // Promoted field
}
