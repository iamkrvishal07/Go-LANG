package main

import "fmt"

const pokemon int = 7

func main() {

	const name string = "GoLang"

	//name := "javastring" // cannot assign

	// const ko function ke bahar v declare kr skte
	fmt.Print(pokemon)
	// := shorthand wla ko func ke bahar declare nhi kr skte


	const(
		PORT = 5;
		HOST = "STRING"
	)
}