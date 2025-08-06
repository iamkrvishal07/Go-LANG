package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict

func main() {
	// creating map

	// naam:=make(map[key]value)
	m := make(map[string]string)

	// setting an element
	m["fruit"] = "Apple"
	m["copy"] =  "Maths"
	m["phone"] = "nokia"

	// get an element
	fmt.Println(m["copy"])

	// IMP: if key does not exists in the map then it returns zero value
	//fmt.Println(m["xyz"])  -> empty for string , 0 for int , false for boolean

	delete(m, "copy")
	// map aur key do delete
	fmt.Println(m)

	// clear(m)
	// map clear ho jata he



	// aonther way to declare map
	// when we know elements otherwise go for make
	// Quiki make mei declare kr dete hai baad mei declare krte hai

	abc := map[string]int{"a": 1, "b": 2}

	//map boolean return krta he

	val , rise := abc["a"] // map-ka-naam[key]

	// val - value 
	// rise - boolean value

	fmt.Println(val)

	if rise {
		fmt.Println("it available")
	} else {
		fmt.Println("Not available")
	}
	fmt.Println(abc)


	abc1 := map[string]int{"a": 1, "b": 2}
	abc2 := map[string]int{"a": 1, "b": 2}

	fmt.Println(maps.Equal(abc1,abc2))



}
