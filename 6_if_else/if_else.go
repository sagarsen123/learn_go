package main

import "fmt"

func main() {
	// if else condition in go
	num := 10
	if num < 18 {
		fmt.Println("You are not eligible to vote")
	}

	// if else condition in go
	if num < 12 {
		fmt.Println("You are a child")
	} else {
		fmt.Println("You may be an adult or teenager")
	}

	// if else if else condition in go
	num = 18
	if num < 12 {
		fmt.Println("You are a child")
	} else if num >= 12 && num < 18 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are an adult")
	}

	// nested if else condition in go
	num = 25
	if num > 12 {
		if num > 20 {
			fmt.Println("Num is grater than 20")
		} else {
			fmt.Println("Num is between 12 and 20")
		}
	} else {
		if num > 5 {
			fmt.Println("Num is between 5 and 12")
		} else {
			fmt.Println("Num is less than 5")
		}
	}

	// conditional statement with short statement
	role := "admin"
	hasPermision := false

	if role == "admin" || hasPermision {
		fmt.Println("You are allowed to access the admin panel")
	}

	role = "user"
	hasPermision = true // check with false also

	if role == "user" && hasPermision {
		fmt.Println("You are allowed to access the user panel")
	} else {
		fmt.Println("You are not allowed to access the user panel")
	}

	// if with short statement
	// if declare var; condition {}
	// this var can be accessed if else block
	// multiple var can be declared using comma
	if score, vinay := 75, 56; score >= 90 {
		fmt.Println("You got A grade")
	} else if score >= 75 {
		fmt.Println("You got B grade")
		print(vinay)
	} else if score >= 60 {
		fmt.Println("You got C grade")
	} else {
		fmt.Println("You got D grade")
	}

	// go does not support ternary operator
	// var result string = condition ? value1 : value2
	// instead we can use if else
	var result string
	marks := 40
	if marks >= 33 {
		result = "Pass"
	} else {
		result = "Fail"
	}
	fmt.Println("Result is", result)
}
