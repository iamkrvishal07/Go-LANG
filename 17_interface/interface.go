package main

import "fmt"

type payment struct{
	//gateway1 stripe
	//gateway2 razorpay

	gateway paymenter

}

type paymenter interface{
	pay(amount float32)
}
// NOTE: interface ke andar jo method he nah uske sign like parameter type same bnayenge kisi struct ke method ke andaar
// toh automatically 

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	//p.gateway1.pay1(amount)
	// p.gateway2.pay2(amount)

	p.gateway.pay(500)
} 

// -> stripe add krne ke baad code modify krna pda jo dikkat ka baat he
// -> Open close principle

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorPay",amount)
}

// func (r razorpay) pay2(amount float32) {
// 	// logic to make payment
// 	fmt.Println("making payment using razorPay",amount)
// }

// type stripe struct{}

// func (s stripe) pay1(amount float32){
// 	fmt.Println("making payement using stripe",amount)
// }


//-----------------------------------------------

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal",amount)
}

func main() {

	// razorpayPaymentGw := razorpay{}
	// newPayment := payment{
	// 	gateway2: razorpayPaymentGw,
	// }

	// newPayment := payment{}
	// newPayment.makePayment(20)

	//interface-GateWAY

	// razor:= razorpay{}
	paypal := paypal{}

	newPayment := payment{
		//gateway: razor,
		gateway: paypal,
	}
	newPayment.makePayment(100)


}