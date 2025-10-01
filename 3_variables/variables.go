package main

import "fmt"

func main() {
	// Declare a variable with a specific type
	var name string = "golang"
	fmt.Println(name)

	// dynamic asigning
	var lastname = "is awesome"
	fmt.Println(lastname)

	var lastnamae = 12
	fmt.Println(lastnamae)

	// int
	var age int = 30
	fmt.Println(age)

	// Short hand declaration
	pk := "31"
	num := 1234
	fmt.Println(pk)
	fmt.Println(num)

	age = 23
	fmt.Println(age)

	// early declaration but later assignment
	// here type of var is neccessary to declare first if we are going to assign it later on
	var city string
	city = "Jabalpur"
	fmt.Println(city)

	var price float32
	price = 12.45
	fmt.Println(price)

	var price2 float64 = 123.4567890123456789
	fmt.Println(price2)

	price3 := 123.4567890123456789
	fmt.Println(price3)

	// boolean

	var isCool bool = true
	fmt.Println(isCool)
}
