package main

import "fmt"

// number sequence of specific length is known as array
// array is a collection of elements of same datatype
// array is a fixed length data structure
// once an array is created its length cannot be changed
// arrays are value types in go
// arrays are indexed starting from 0
func main() {

	// delaring an array
	// type name_of_array [length]datatype
	// by default each value will be assigned to the default value of its datatype here 0 for int

	var nums [5]int
	fmt.Println(nums, len(nums)) // len function returns the length of the array

	// assigning values to array
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50
	// nums[5] = 60 // this will throw an error as index out of range

	// accessing values from array
	fmt.Println("First element:", nums[0])
	fmt.Println("Second element:", nums[1])
	fmt.Println("Third element:", nums[2])
	fmt.Println("Fourth element:", nums[3])
	fmt.Println("Fifth element:", nums[4])

	var boolarr [3]bool
	fmt.Println(boolarr) // default value of bool is false

	boolarr[0] = true
	boolarr[1] = false
	boolarr[2] = true
	fmt.Println(boolarr)

	// float array
	var floatarr [4]float64
	fmt.Println(floatarr)
	floatarr[0] = 1.1
	fmt.Println(floatarr)

	var floatarr2 [4]float32
	fmt.Println(floatarr2)
	floatarr2[0] = 1.1
	fmt.Println(floatarr2)

	// string array
	var strarr [3]string
	fmt.Println(strarr) // default value of string is empty string
	strarr[0] = "hello"
	fmt.Println(strarr)

	// array literal
	// array can be initialized at the time of declaration using array literal
	// array literal is a comma separated list of values enclosed in curly braces
	// the length of the array is determined by the number of values in the array literal
	// syntax: name_of_array := [length]datatype{value1, value2, value3, ...}
	// if length is not specified then it will be determined by the number of values in the array literal
	// syntax: name_of_array := [...]datatype{value1, value2, value3, ...}
	numsArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numsArr)

	strArr := [...]string{"go", "is", "awesome"}
	fmt.Println(strArr)
	fmt.Println("Length of strArr:", len(strArr))

	// iterating over an array
	// using for loop
	for i := 0; i < len(strArr); i++ {
		fmt.Println("Element at index", i, "is", strArr[i])
	}

	// 2d arrays in go
	var matrix [2][3]int
	fmt.Println(matrix)
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	fmt.Println(matrix)

	// 2d array literal
	// syntax: name_of_array := [rows][cols]datatype{{value1, value2, ...}, {value1, value2, ...}, ...}
	var matrix2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrix2)

	// slices in array
	// slices are used as dynamic arrays in go
	// use arrays if the size is fixed and known at compile time
	// slices are a more powerful and flexible way to work with sequences of data than arrays
	// slices are built on top of arrays and provide a dynamic view into the elements of an array
	// slices are reference types in go
	// slices do not store any data, they just point to an underlying array
	// slices can be resized and can grow and shrink as needed
	// slices have a length and a capacity
	// length is the number of elements in the slice
	// capacity is the number of elements in the underlying array that the slice can access
}
