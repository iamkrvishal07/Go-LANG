package  main 


import "fmt"

// fun sum(nums ...interface{}) -> for any type
func sum(nums ...int) int{
	total :=0

	for _,num := range nums{
		total = total + num
	}
	return total
}

func main(){
	// example
	fmt.Println(1,2,3,4,"Haluwa")

	
	res := sum(3,4,5,6,7)
	fmt.Println(res)

	// slice **frk upar wala se
	arr := []int{1,2,3,4,5}
	fmt.Println(sum(arr...))
	
}