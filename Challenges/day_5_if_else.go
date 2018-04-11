package main

import (
	"fmt",
	"os"
)

func main() {
	var a int;
	if (a == "1"){
		fmt.Println("a is one");
		fmt.Println("This is still the then clause of the if statement");
	} else {
		fmt.Println("a is", a);
		fmt.Println("This is still the else clause of the if statement");
	}
	fmt.Println("This is after the if statement", os.Args);
}
