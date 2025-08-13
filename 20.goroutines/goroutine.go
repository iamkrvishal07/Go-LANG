package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("TASK IS DONE BY IDno - ",id)
}

func main() {

	for i :=1 ; i<=5; i++{
		go task(i)
	}

	time.Sleep(time.Second *2 )

}