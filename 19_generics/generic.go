package main

import "fmt"

func printSlice[T comparable](items []T){
	for _,item := range items{
		fmt.Println(item)
	}
}

// func printSlice[T int | string](items []T){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T any](items []T){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice(items []int){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSliceString(items []string){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }


//LIFO

type stack[T any] struct {
	elements []T
}

// type stack struct {
// 	elements []int
// }

func main() {
	fmt.Println("hello-world")
	nums := []int{1,2,3}
	printSlice(nums)

	names :=[]string{"golang", "typescript"}
	// printSliceString(names)
	printSlice(names)

	myStack :=stack[int]{   // -> yaha batana pdta he kis type ka he yeh [int] [string]
		elements: []int{1,3,4},
	}
	fmt.Println(myStack)
}