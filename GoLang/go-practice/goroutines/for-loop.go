package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)


func Test_Go_Routines(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("i value is %v\n", n)
		amt := time.Duration(rand.Intn(2000))
		time.Sleep(time.Millisecond * amt)
	}
}
