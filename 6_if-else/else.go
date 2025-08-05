package main

import "fmt"

func main() {

	// age := 18

	// if age >= 19 {
	// 	fmt.Println("Person is an adult")
	// } else{
	// 	fmt.Println("Person is not an adult")

	var role = "admin"
	var hasPermissions =true

	if role =="admin" || hasPermissions{
		fmt.Println("yes")
	}

	if role =="admin" && hasPermissions{
		fmt.Println("yes")
	}

	// we can declare a variable inside if contruct
	if age:= 25;age>=16{
		fmt.Println("person is boy")
	}

	// Go does not have ternary , u will have to use normal if else


	
}