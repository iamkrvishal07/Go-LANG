package main

import "fmt"

// by value // distinct value // copy

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum",num)
// }

// by reference

func changeNum(num *int) {
	// dereference
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	// changeNum(num)
	fmt.Println("Before changeNum in main", num)
	changeNum(&num) // address beja maine
	fmt.Println("After changeNum in main", num)
	// output - InchangeNum 5 , After changeNum in main 1
}
