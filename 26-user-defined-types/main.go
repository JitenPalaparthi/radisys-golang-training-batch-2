package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	var mymap1 myMap

	mymap1 = make(myMap) // creates 0 length
	//mymap = make(map[string]any, 10)
	var map1 map[string]Object
	fmt.Println(mymap1, len(mymap1)) // 0 length
	map1 = map[string]Object(mymap1) // it works without this
	mymap1["Bangalore-1"] = 560086
	mymap1["Bangalore-2"] = 560034
	mymap1["Trivandrum-1"] = 690031
	mymap1["Trivandrum-2"] = 690038
	fmt.Println(map1, len(map1)) // 0 length

	map1["Hyderabad-1"] = "540043"
	fmt.Println(mymap1, "type of mymap1", reflect.TypeOf(mymap1))

	if ok, err := mymap1.Delete("Hyderabad-1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Key has been delete:", ok)
	}

	if ok, err := mymap1.Delete("Hyderabad-1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Key has been delete:", ok)
	}

	if ok, err := myMap(map1).Delete("Bangalore-2"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Key has been deleted:", ok)
	}
	fmt.Println(mymap1, "type of mymap1", reflect.TypeOf(mymap1))
}

type myMap map[string]Object

func (mm myMap) Delete(key string) (bool, error) {

	if mm == nil {
		return false, errors.New("nil map")
	}

	if _, ok := mm[key]; !ok {
		return false, errors.New("no key")
	} else {
		delete(mm, key)
		return true, nil
	}
}

// only used as alias
//type Object = any

type Object = any

// implement all these methods on myMap--> insert, update, get , getKeys()[]string
// getValues()[]any
// all methods must have error as one of the return types. nil or empty to be implemented.
