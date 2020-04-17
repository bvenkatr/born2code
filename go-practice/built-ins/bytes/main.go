package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Interpret Comapare's result by comparing it to zero
	var a, b []byte
	a = []byte("a")
	b = []byte("b")
	if bytes.Compare(a, b) < 0 {
		fmt.Println("a less than b")
	}

	if !bytes.Equal(a, b) {
		fmt.Println("a is not equal to b")
	}

	fmt.Println(bytes.Count([]byte("Chse"), []byte("e")))
}