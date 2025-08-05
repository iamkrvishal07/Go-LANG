package main

import "fmt"

func main() {
	// var--variable ka naam---variable ka type

	var name string = "Hello-GoLang"
	fmt.Println(name)

	// infer
	var math  = 1 
	fmt.Println(math)

	//shorthand syntax
	a:=1
	b:=2
	fmt.Println(a+b);


	// baad mei asign krna he toh

	var pokemon string
	// yaha bs variable define he 
	// hum chahte ki baad mei use kre isliye

	pokemon="Pikachu"

	fmt.Println(pokemon)

	// var price float32 = 100.5
	// var price = 100.5
	price:=100.5

	fmt.Println(price)


}