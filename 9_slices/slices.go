package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in Go
// + useful methods

func main() {
	// uninitialized slice is nil
	var nums []int

	fmt.Println(nums)
	// output - []
	// [] == nil => true

	fmt.Println(len(nums))
	// len = 0;

	var arr = make([]int, 2)
	fmt.Println(arr)
	// [0,0]
	// now nums==nil -> false

	var num1 = make([]int, 2, 5)
	// capacity -> maximum numbers of elements can fit

	num1 = append(num1, 1)
	num1 = append(num1, 1)
	num1 = append(num1, 1)
	num1 = append(num1, 1)
	num1 = append(num1, 1)

	fmt.Println(cap(num1)) // 10
	fmt.Println(num1)      // [ 0 0 1 1 1 1 1]
	// for normal slices --> var num1 = make([]int, 0, 5)
	// cap now becomes 2x


	//var num1 = make([]int, 0, 5)
	// num1[0] = 2 // error initalise length 0 so , make it initial 2
	// then num1[0]=1
	// num2[1]=2



	// another way

	poke := []int{}
	poke = append(poke, 1)
	poke = append(poke, 2)
	poke = append(poke, 3)
	poke = append(poke, 4)
	poke = append(poke, 5)

	fmt.Println(poke)
	fmt.Println(cap(poke)) // 2 4 8 16 jta he
	fmt.Println(len(poke))


	// copy function

	var cp1 = make([]int,0,5)
	cp1=append(cp1, 1)
	var cp2 = make([]int, len(cp1))

	fmt.Println(cp1,cp2)  // [1][0]
	copy(cp2,cp1) // kaha se kisse
	fmt.Println(cp1,cp2) // [1][1]


	// Slice operator

	var num3= []int{1,2,3}
	fmt.Println(num3[0:2]) // 1,2
	fmt.Println(num3[:2])  // 1,2
	fmt.Println(num3[1:]) //  2,3


	// slice package

	var z1 = []int{1,2}
	var z2 = []int{3,4}

	fmt.Println(slices.Equal(z1,z2)) // return boolean -> here False

    // for 2D array
	var array2d = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(array2d)

}
