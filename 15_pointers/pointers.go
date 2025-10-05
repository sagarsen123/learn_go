package main

import "fmt"

// here the arguments are passed by value so value will remain unchange in main
func changeNum(num int) {
	num = 5
	fmt.Println(num)
}

// passing by reference
// accepting a pointer
func changeNumRef(num *int) {
	// modifying the value in num
	// &x of var x returns its address
	// *x return of a pointer returns value of a number
	*num = 7
	fmt.Println(*num)
}

func main() {

	// pointers stores the memory addresses of the variables
	// lets take example of num
	num := 1

	// try to change value of no
	// but they are passed by value
	changeNum(num)
	fmt.Println(num)

	// tring again with pass by reference
	// using pointers
	changeNumRef(&num) //passing memory address
	fmt.Println(num)
}
