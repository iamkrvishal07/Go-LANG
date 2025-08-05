package main

import "fmt"

func main() {

	//QuickNote-->
	// int ->0, string ->"", bool->false

	var arr [4]int
	arr[0]=1

	fmt.Println(len(arr))
	fmt.Println(arr[0])
	// whole array print
	fmt.Println(arr)
	// output -- [1,0,0,0]


	// for boolean arr
	var vals [4]bool
	fmt.Println(vals)
	// [false false false false]

	var naam [4]string
	fmt.Println(naam)
	// empty string



	// to declare arr in sinle line
	nums:= [4]int{1,2,3,4}
	fmt.Println(nums)


	// 2d arrays

	nums1:= [2][2]int{{3,4},{5,6}}
	fmt.Println(nums1)


	// -fixed size, that is predictable
	// -Memory optimization
	// -contant time access

}