package main

import "fmt"

// syntac
// func funciion_name ( paramter_name parameter_type, ....) return type { function body}
func add(a int, b int) int {
	return a + b
}

// here the syntac is same but we can skips mulitple data type decleration for parameter of same type
//  we can write parameters as a,b,c,d ... parameter_type.
// This will make all tha parameter before type of the same kind
func multi(a, b int) int {
	return a * b
}

// go functions can return multiple values
//  we need to define type of the each returning values
func getLang() (string, string, string) {
	return "Go", "Js", "C"
}

func giveNo(fn func(a int) int) {
	fn(5)
}

func getFunction(op int) func(a, b int) int {
	add := func(a, b int) int {
		return a + b
	}
	subs := func(a, b int) int {
		return a - b
	}
	multi := func(a, b int) int {
		return a * b
	}
	div := func(a, b int) int {
		return a / b
	}
	switch op {
	case 1:
		return add
	case 2:
		return subs
	case 3:
		return multi
	case 4:
		return div
	default:
		return func(a, b int) int {
			return 0
		}
	}
}

func main() {
	fmt.Println(add(10, 15))
	fmt.Println(multi(5, 10))
	a, b, c := getLang()
	fmt.Println(a, b, c)

	d, e, _ := getLang()
	fmt.Println(d, e)

	// passing funcitons as parameter
	fn := func(a int) int {
		fmt.Println(a)
		return a
	}
	giveNo(fn)

	// get function
	received_func := getFunction(3)
	print(received_func(9, 10))
}
