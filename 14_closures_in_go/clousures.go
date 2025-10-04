package main

// Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

import "fmt"

func closure() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	// they both return 1
	// here two different instances of closure are calling different increments
	fmt.Println(closure()())
	fmt.Println(closure()())

	// but if we create an instance
	// this create an instance of clousure returning functions
	// ans same instance will be called to increment returning 1,2,3
	increment := closure()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
