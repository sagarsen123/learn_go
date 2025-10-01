package main

import "fmt"

// externally or globally constants and variables can be declared
//  but the shorthand declaration is not allowed
// format for var declaration
// [type var or const] [name of var ] [datatype of var ] = [value]
const language string = "Go"

var version = 1.19

// apple:="fruit" // this will throw an error
// const applie string

// in go constants should be initialized at the time of declaration
// and their value cannot be changed later
// fruit is a constant with no value assigned, it will through an compilation error, but if we leave the var uninitialized then its value will be default value of it datatype like
// 0 for int
// empty string for string

var fruit string

func main() {
	// Constants are immutable values which are known at compile time and do not change for the life of the program.
	const pi float64 = 3.14
	fmt.Println(pi)

	const age = 12
	fmt.Println(age)

	fmt.Println(language)
	fmt.Println(version)

	// fruit = "apple"
	fmt.Println(fruit)

	// multi constant declaration
	// multiple constant and variables can be declared enclosing a pair of parenthesis
	const (
		port     = 8080
		host     = "localhost"
		protocol = "http"
	)
	fmt.Println(port, host, protocol)

	var (
		name    string = "Golang"
		latency int    = 123
		isGood  bool   = true
	)
	fmt.Println(name, latency, isGood)
}
