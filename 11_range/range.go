package main

import "fmt"

// iterating over data structures

func main(){
	nums :=[]int{4,5,6}


	// for i:= 0 ;i<len(nums);i++ {
	// 	fmt.Println(nums[i])
	// }

	sum:=0;

	//for index , num:= range nums
	for idx,num:= range nums{
		sum+=num
		fmt.Println(sum)
		fmt.Println(idx)
	}

	// for MAP

	m:= map[string]string{"fanme":"john", "lname":"doe"}

	// key , value := 
	// agar bs ek variable denge toh byDefault key hi consider hota he
	for k,v := range m{
		fmt.Println(k,v)
	}


	// unicode code point rune
	// starting byte of rune
	for i,c := range "goLang¥ÿ" {
		fmt.Println(i,c)
	}

	// output --    // 0 103
					// 1 111
					// 2 76
					// 3 97
					// 4 110
					// 5 103
					// 6 165
					// 8 255 ---> takes 2 byte and increased index 7 to 8


	  

}