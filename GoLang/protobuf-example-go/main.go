package main

import (
	"fmt"

	simplepb "./src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
	sm.Name = "I renamed you"
	fmt.Println(sm.GetId())
}
