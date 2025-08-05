package main

import (
	"fmt"
	//"time"
)

func main() {
	// simple switch
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// default:
	// 	fmt.Println("Other")
	// }



	// multiple condition switch

	// switch time.Now().Weekday(){
	// case time.Saturday, time.Tuesday:
	// 	fmt.Println("It's weekend")

	// default:
	// 	fmt.Println("It's workday")

	// }.


	// type switch

	functionA := func(i interface{}){
		switch t:= i.(type){
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("other",t)
		}
	}

	functionA(1)


	
}