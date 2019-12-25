package main

import (
	test "./formatting"
	"fmt"
)

func main() {
	// test.OutputToStdout()
	filename, err := test.OutputToWriter()
	fmt.Println(filename, err)
}
