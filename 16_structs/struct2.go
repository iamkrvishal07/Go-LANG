package main

import (
	"fmt"
	//"time"
)

// order struct

type customer struct{
	name string
	phone int
}

type order struct{
	id string
	amount int
	status string
	//createdAt time.Time
	customer // ----------------> bas yhi krna tha simple
}
func main() {

	// newCustomer := customer{
	// 	name:"vishal",
	// 	phone:1234567890,
	// }

	myOrder :=order{
		id: "1",
		amount :70,
		status: "true",
		//customer:newCustomer,
		customer: customer{
			name:"vishal",
			phone: 91929191,
		},
	}

	fmt.Println(myOrder)


	
}