package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch case in go
	//  we dont require to add break statement in each case as in other languages like c, java etc
	//  because by default go breaks the switch after executing a case
	//  if we want to continue the execution of next case then we can use fallthrough statement
	//  but its not recommended to use fallthrough statement as it can lead to unexpected behavior
	//  its better to use if else if else ladder in such cases
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// switch with multiple cases in a single case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// type switch in go
	// interface is a type that can hold any value
	whoamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a boolean")
		default:
			fmt.Printf("I am of a different type %T\n", t)
		}
	}
	whoamI(12)
	whoamI("hello")
	whoamI(true)
	whoamI(12.34)

	// simplefied
	// variable can also be removed from switch statement
	// and we can directly use the type assertion in switch cases
	// this is useful when we dont need to use the variable in the switch statement
	findwhoamI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a boolean")
		default:
			fmt.Printf("other")
		}
	}

	findwhoamI(12)
	findwhoamI("hello")
	findwhoamI(true)
	findwhoamI(12.34)
}
