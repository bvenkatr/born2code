package main

import (
	"fmt"

	example_simple "./src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
}
