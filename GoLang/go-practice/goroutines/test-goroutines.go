package main

import (
	"fmt"
)

func main() {
	var a chan int
	if a == nil {
		fmt.Println("Channel is nil, so defining it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
}
