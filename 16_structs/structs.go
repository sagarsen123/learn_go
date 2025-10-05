package main

import (
	"fmt"
	"time"
)

// structs are building block of oops programming in go
// with structs we can create custom datatype

type order struct {
	id        string
	amount    float32
	status    string
	quantity  int
	createdAt time.Time //it has nano second precision
}

// constructs with structs
// go directly doesn't provide constructors
// we use a method to work like constructors
// it is convention to have the name started with 'new' then in camel case the type of struct
// like constructor for order will have newOrder name
func newOrder(id string, amount float32, status string) *order {
	// we create an object of struct here
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	// here we return the pointer to the struct object
	return &myOrder
}

// we attach method/behaviours to a struct
// take an example of status changing method
// receiver types are used to connect methods to the structs
// there is a convention while creating the receiver types, we need to name of receiver var with the starting char of struct type
// as o for order type then define the order

// we need to pass the variable by reference if we want to change values
//passing by value will not change anything like here
// func (o order) changeStatus(status string) {
// 	o.status = status
// }

// this will change the status
// struct automatically performs dereferencing
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func main() {
	// creating instance of a struct
	// we can add as many fields as we wanted in the order.
	// it will not through an error even if we leave the createdAt field unfill
	// it will automatically assign a default value to it
	// as int to 0
	// datetime to nill value
	var od order = order{
		id:     "1",
		amount: 100.00,
		status: "received",
	}

	// shorthand
	my_order := order{
		id:     "2",
		amount: 50.00,
		status: "received",
	}
	fmt.Println(od, my_order)

	// empty fields can be added later on
	// we can access and assign values to struct fields using '.' dot operator
	my_order.createdAt = time.Now()

	fmt.Println(my_order, my_order.amount, my_order.id)

	// updating fields using struct methods
	od.changeStatus("Paid")
	fmt.Println(od)
	fmt.Println("Order Amount", od.getAmount())

	// creation using constructor
	bagOrder := newOrder("3", 400, "placed")
	fmt.Println("bag order", bagOrder, bagOrder.id)

	// in-line struct creation
	language := struct {
		name   string
		isGood bool
	}{
		"golang", true,
	}

	fmt.Println(language)

}
