package main

import (
	"fmt"
	"maps"
)

// hash, object, dictionary
func main() {
	// Create a map with string keys and int values
	// syntax: name_of_map = make(map[key_type]value_type)
	ages := make(map[string]int)

	// add value to the map
	ages["Alice"] = 30
	ages["Bob"] = 25

	// get elemeent from map
	fmt.Println("Alice's age is", ages["Alice"])

	// get non existing element
	fmt.Println("Eve's age is", ages["Eve"]) // returns 0, the zero value for int

	// length of map
	fmt.Println("Number of entries in ages map:", len(ages))

	// check if key exists
	age, ok := ages["Bob"]

	if ok {
		fmt.Println("Bob's age is", age)
	} else {
		fmt.Println("Bob's age not found")
	}
	fmt.Println(ages)
	_, isok := ages["Eve"]
	if isok {
		fmt.Println("every this in ok")
	} else {
		fmt.Println("Bob's age not found")
	}
	// clear map
	clear(ages)

	// create an initialized map
	m := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println(m)

	// delete element from map
	delete(m, "foo")
	fmt.Println(m)

	// iterate over map
	for key, value := range m {
		fmt.Println(key, value)
	}
	// iterate over map keys
	for key := range m {
		fmt.Println("key:", key)
	}
	// iterate over map values
	for _, value := range m {
		fmt.Println("value:", value)
	}
	// using maps package to
	// get keys of map
	keys := maps.Keys(m)
	fmt.Println("Keys:", keys)

	// get values of map
	values := maps.Values(m)
	fmt.Println("Values:", values)

	// check maps are equal
	m1 := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	m2 := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	fmt.Println(maps.Equal(m1, m2))
}
