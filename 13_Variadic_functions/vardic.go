package main

import "fmt"

// variadic funtions are the functions that can have any no of arguments/parameters as an input

// our custom variadic funtion
// syntax : func funtion_name(parameter_name ...parameter_Datatype) returntype { function body}
func sum(num ...int) int {
	total := 0
	for _, i := range num {
		total += i
	}
	return total
}

func main() {
	// this is a variadic funtions
	fmt.Println(1, 2, 3, 4, "Hello")
	fmt.Println(sum(1, 2, 3, 4, 5))

	// consider data in slice
	nums := []int{10, 20, 30, 40, 50}
	sum_nums := sum(nums...)
	fmt.Println("Sum of nums slice", sum_nums)
}
