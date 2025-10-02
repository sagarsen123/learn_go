package main

import "fmt"

func main() {
	fmt.Println("Range in Go")

	// iterate over array
	arr := []int{10, 20, 30, 40, 50}
	sum := 0
	for _, value := range arr {
		fmt.Println("Values", value)
		sum += value
	}
	fmt.Println("Sum of array elements:", sum)

	// rnage over map
	m := map[string]string{
		"1": "John",
		"2": "Vinay",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// iterate over string
	for i, c := range "This is a go string" {
		// here i is starting byte of the rune
		//  here c is unicode value of char c
		fmt.Println(i, c)
	}

	// get the char value in string using loop
	for i, c := range "This is a go string" {
		//string(c) this will return the real value of c
		fmt.Println(i, string(c))
	}

}
