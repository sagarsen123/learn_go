package main

import "fmt"

// part 1
// payment struct
// type payment struct{}

// func (p payment) makePayment(amount float32) {
// 	// lets add a payment gateway
// 	// razorPayGw := razorPay{}
// 	// razorPayGw.pay(amount)

// 	// suppose if we have to add stripe then we have to modify and remove razorpay
// 	stripeGW := stripe{}
// 	stripeGW.pay(amount)

// 	// we can not make this type of implemnetation in real world
// 	// suppose if we have add another gateway this will not be a good practice to modify our makePayment struct
// }

// // part 2
// // suppose in order to resolve this we can add a gateway to out payment stuct and can make payment
// type payment struct {
// 	gateway stripe

// 	// but this doesn't resolve our problem, we have hardcoded the gateway type here, so if we have to add or change the gate way we again have to modify the payment.
// 	// we need a general type so that no of gateways can be added removed without affecting the main functionality
// 	// here the interfaces comes into the picture
// }
// func (p payment) makePayment(amount float32) {
// 	p.gateway.pay(amount)
// }

// part 3
// interfaces implementations
// there is naming convention for the interfaces like we want to make interface for the payment, so we have to name it as paymenter, 'suffix' is added to the payment
// syntax: type name_of_interface interface{ interface_methods.. }
type paymenter interface {
	// here we define the interface methods
	// these interface methods are automatically detected by go in structs.
	//  we dont have to mention explicitly about who implements this interface.
	pay(amount float32) // syntax : method_name(params paramstype) return type
	// a no of method can be added to a interface
	// struct should implement all methods metioned in the interface with same signatures or else it will not be recognized as a implementer of interface
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

// creating gateway
type razorPay struct{}

func (r razorPay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

// creating stripe gateway
type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)
}

func main() {
	fmt.Println("Interfaces in go lang")

	// part 1
	// creating payment struct instance
	// newPayment := payment{}
	// newPayment.makePayment(100)

	// // part 2
	// // now we have to pass the gw explicitly from here
	// stripeGW := stripe{}
	// newPayment := payment{stripeGW}
	// newPayment.makePayment(100)

	// // this will not work since we have stipe gateway hardcoded
	// razorPayGW := razorPay{}
	// newPayment2 := payment{razorPayGw}
	// newPayment2.makePayment(100)

	// part 3
	// interfaces implementation
	razorPayGw := razorPay{}
	stripeGw := stripe{}
	new_payment := payment{gateway: razorPayGw}
	new_payment.makePayment(100)

	// now we can simply create or pass the gateway add or remove without affecting the main functionality
	new_payment2 := payment{gateway: stripeGw}
	new_payment2.makePayment(100)
}
