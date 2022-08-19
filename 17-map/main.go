package main

import "fmt"

// what can be a key? Any type where you can implement == operator can be a key
// can be implemented string, int, float, bool
// cannot be implemented on any, struct or some other types
// can check whether a map is nil or not
// to instantiate a map use make function
// map is a reference type
func main() {
	var mymap map[string]string // this is just a declaration but not instantiation

	if mymap == nil {
		fmt.Println("mymap is nil")
	}

	mymap = make(map[string]string)

	mymap["Bangalore-1"] = "560086"
	mymap["Bangalore-2"] = "560096"
	mymap["Trivandrum-1"] = "590034"
	mymap["Hyderabad-1"] = "544678"
	mymap["Chennai-1"] = "544034"

	fmt.Println(mymap)

	for k, v := range mymap {
		fmt.Println("Key:", k, "Value:", v)
	}

	val := mymap["Bangalore-2"]

	fmt.Println("Value:", val)

	val, ok := mymap["Bangalore-1"]
	fmt.Println(val, ok)

	val, ok = mymap["Bangalore-3"]
	fmt.Println(val, ok)

	delete(mymap, "Chennai-1")
	for k, v := range mymap {
		fmt.Println("Key:", k, "Value:", v)
	}
}

// Create slice or arrry with few elements
// find duplicate elements in an array using map

// update map--> if no key exist/nil return an error "no key exist so not updated."

// delete map --> if no key exist/nil return an error "no key or nil"
