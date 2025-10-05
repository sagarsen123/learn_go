package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    int
	status    string
	createdAt time.Time

	// a reference to another struct can also be added into a struct
	buyer customer
}

// nested structs or embeddings can be used to implement inheritance and composition in go

func main() {
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		buyer: customer{
			name:  "John",
			phone: "1234467894",
		},
	}
	fmt.Println(newOrder)

	newCustomer := customer{
		name:  "Ryan",
		phone: "7788994455",
	}
	newOrder.buyer = newCustomer
	fmt.Println(newOrder)

	// nested accessing of variables is done with dot operators
	newOrder.buyer.phone = "8888994455"
	fmt.Println(newOrder)
}
