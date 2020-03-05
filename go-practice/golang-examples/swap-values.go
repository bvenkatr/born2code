package main

import "fmt"

func main() {
	fmt.Println("test")

	// Swap two values
	x, y := 1, 2
	fmt.Println("x \t y")
	fmt.Println(x, "\t", y, "before")
	x, y = y, x
	fmt.Println(x, "\t", y, "after")
}
