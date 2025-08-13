package main

import "fmt"

// enumerated types
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status string) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus("Received")
	fmt.Println("Delivered status value is", Delivered)
}
