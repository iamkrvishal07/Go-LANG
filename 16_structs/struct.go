package main

import (
	"fmt"
	"time"
)

//order struct <blueprint>

type order struct{
	id     string
	amount float32
	status string
	createdAt time.Time   // time package inbuilt -> nanosecond precision
}


// receiver type
func (o *order)changeStatus(status string){
	o.status=status
	// struct automatically do deferencing work.....

}

func (o *order)getAmout() float32{
	return o.amount
	// yaha * use kr v skte the nhi v [ agar value change krna he toh * use kro best ]
	// yaha (o order) v likhe toh kaam karega .
}



func newOrder(id string, amount float32, status string)*order {
	// initial setup goes here...
	myOrder :=order{
		id: id,
		amount: amount,
		status: status,
	}

	return &myOrder


}




func main() {


	myOrder := newOrder("1",30.50,"received")
	fmt.Println(myOrder)

	// ** if you don't set any field, default value is zero value
	// int => 0, float=>0, string "", bool => false

	// var <name> order =

	//shortHand Syntax

	// myOrder := order{
	// 	id: "1",
	// 	amount: 70.00,
	// 	status: "pending",
	// }

	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.amount)
	// fmt.Println("Order struct", myOrder)


	// 	myOrder := order{
	// 	id: "1",
	// 	amount: 70.00,
	// 	status: "pending",
	// }


		myOrder := order{
		id: "1",
		amount: 70.00,
		status: "pending",
	}

	//myOrder.status = "rejected"
	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder)
	fmt.Println(myOrder.getAmout())

}