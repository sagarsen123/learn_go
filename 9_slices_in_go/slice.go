package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice is a dynamic array
	// slice is a reference type
	fmt.Println("Slices in Go")

	// slices are the dynamic version of arrays
	// slices do not have a fixed length
	// slices are built on top of arrays
	// slices are more flexible than arrays
	// slices are more powerful than arrays
	// slices are more efficient than arrays
	// slices provide many built-in functions to work with them

	// declaring a slice
	// syntax: var name_of_slice []datatype
	// uninitialized slice will have a default value of nil
	// length and capacity of an uninitialized slice will be 0
	// length is the number of elements in the slice
	// capacity is the number of elements in the underlying array
	var nums []int
	fmt.Println(nums) // default value of slice is nil
	fmt.Println("Length of nums slice:", len(nums))
	fmt.Println("Capacity of nums slice:", cap(nums))
	fmt.Println(nums == nil) // true

	// initializing a slice
	// using make function
	// syntax: name_of_slice := make([]datatype, inital_length, initial_capacity)
	// if capacity is not specified then it will be equal to length
	nums = make([]int, 5)
	fmt.Println(nums) // here defullt values of int is 0
	fmt.Println("Length of nums slice:", len(nums))
	fmt.Println("Capacity of nums slice:", cap(nums)) //max no of elements that can be stored in the slice

	// create empty slice with 0 length and 5 capacity
	emptySlice := make([]int, 0, 5)
	fmt.Println(emptySlice)
	fmt.Println("Length of emptySlice:", len(emptySlice))
	fmt.Println("Capacity of emptySlice:", cap(emptySlice))

	// adding values to the slice
	nums = append(nums, 1)
	nums = append(nums, 2, 3, 4, 5)
	fmt.Println(nums)
	fmt.Println("Length of nums slice:", len(nums))
	fmt.Println("Capacity of nums slice:", cap(nums)) // capacity will be doubled when the length exceeds the capacity

	// accessing elements of a slice
	fmt.Println("Element at index 0:", nums[0])

	// modifying elements of a slice
	nums[0] = 10
	fmt.Println("Modified slice:", nums)

	// syntac for appending another slice to the slice
	// append(slice1, slice2...) // note the ... at the end of slice2
	nums = append(nums, []int{6, 7, 8}...) // appending another slice to the slice
	fmt.Println(nums)
	fmt.Println("Length of nums slice:", len(nums))
	fmt.Println("Capacity of nums slice:", cap(nums)) // capacity will be doubled when the length exceeds the capacity

	// shorthand declaration of slice
	shortArray := []int{}
	fmt.Println(shortArray)

	sl1 := []int{1}
	sl2 := []int{2, 3, 4}
	fmt.Println(sl1, sl2)

	// case 1
	// create two slices, initialize
	//  here this code will run, but here the length of both the slices is 0
	//  so when we will try to copy the elements from one slice to another then nothing will be copied
	//  because the length of the destination slice is 0
	var numsog = make([]int, 0, 5)
	var numvcp = make([]int, len(numsog), 5)
	numsog = append(numsog, 1, 2, 3)

	copy(numvcp, numsog)

	fmt.Println(numsog, numvcp)

	// now to copy the elements from one slice to another
	// we need to make sure that the length of the destination slice is greater than or equal to the length of the source slice
	// case 2
	var numsog2 = make([]int, 0, 5)
	numsog2 = append(numsog2, 1, 2, 3) // length of numsog2 is 3 now
	// so to copy the elements from numsog2 to numscp2, we need to make sure that the length of numscp2 is at least 3
	// we can do this by using len(numsog2) to set the length of numscp2
	var numscp2 = make([]int, len(numsog2))
	copy(numscp2, numsog2)

	fmt.Println(numsog2, numscp2)

	// slice operator
	ogSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Sliced daa", ogSlice[1:3]) // from index 1 to index 2 excluding 3
	fmt.Println("Sliced daa", ogSlice[:3])  // from index 0 to index 2 excluding 3

	// slice package
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	fmt.Println("Are slices equal?", slices.Equal(slice1, slice2)) // true
	//  many method are present in slices package like
	// slices.Contains, slices.Index, slices.Clone, slices.Delete, slices.Insert, slices.Reverse, slices.Sort, slices.Compact etc.

	// 2d slices in go
	// syntax:  var name_of_slice [][]datatype
	//  here each element of the outer slice is an inner slice
	//  so we can append inner slices to the outer slice
	var matrix [][]int //syntacs for 2d slice
	fmt.Println(matrix)
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	fmt.Println(matrix)
	// accessing elements of 2d slice
	fmt.Println("Element at row 0, column 0:", matrix[0][0])
	fmt.Println("Element at row 1, column 2:", matrix[1][2])

	// initializzing 2d slice
	matrix2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrix2)

}
