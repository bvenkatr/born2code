package strconv_examples

import (
	"fmt"
	"strconv"
)

func StringToInt() {
	n, err := strconv.Atoi("24")
	if err != nil {
		fmt.Println("Got error while doing string to integer conversion ", err)
		return
	}
	fmt.Printf("%v - %T\n", n, n)
}
