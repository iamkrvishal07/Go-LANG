package main

import "fmt"


// also we can write func add(a,b int) 
// --> yeh samj jta he ki aage wala int type ka he

func add(a int, b int) int {
	return a + b
}

// example of multiple return values
func getLang() (string,string,int) {
	return "Golang","apple",25
}


// func return ho rha he 
func processIt() func(a int) int{
	return func(a int) int{
		return 4
	}
}

func main() {

	fn :=processIt()
	fn(7)

	z := add(4, 5)
	fmt.Println(z)
	fmt.Println(getLang())

	// We can also do 
	a,b,_ := getLang() // underscore**

	fmt.Println(a);
	fmt.Println(b)
	// fmt.Println(n)



}