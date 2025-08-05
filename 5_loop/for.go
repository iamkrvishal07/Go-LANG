package main

import "fmt"

// for-> only constuct in go for looping
func main(){

	//while loop
	i:=1
	for i<=5{
		fmt.Println(i)
		i++;
	}

	// infinite loop
	// for {
	// 	fmt.Println(1);
	// }

	// classic for loop

	for i:=0; i<3; i++{
		fmt.Println(i)
	}

	// we can use break || continue 
	for j:=0 ; j<5; j++ {
		if j ==2 {
			continue
		}
		fmt.Println(j)
	}


	// Range ( zero se start hoga )
	for i:= range 3 {
		fmt.Println(i)
	}


}