package main

import "fmt"

func main() {
	// go only provides consttuct for "for" loop

	// standar for loop
	for i := 1; i <= 3; i++ {
		fmt.Println("for loop iteration", i)
	}

	// break and continue in go
	for k := 1; k <= 5; k += 2 {

		fmt.Println("for loop with step iteration", k)

	}
	// break and continue in go
	for k := 1; k <= 5; k += 1 {
		if k == 1 {
			continue
		}

		fmt.Println("for loop with continue and break iteration", k)
		if k == 2 {
			break
		}
	}

	// while loop with for loop
	i := 1
	for i <= 3 {
		fmt.Println("while loop iteration", i)
		i++
	}

	// looping with range
	//  it was introduced in go 1.22
	for i := range 4 {
		fmt.Println("looping with range", i)
	}

	// do while loop with for loop
	j := 1
	for {
		fmt.Println("do while loop iteration", j)
		j++
		if j == 3 {
			break
		}
	}

	//  infinite loop in go suing for
	// j := 1
	// for {
	// 	fmt.Println("infinite loop iteration", j)
	// }

}
