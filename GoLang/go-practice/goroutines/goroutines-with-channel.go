package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello World goroutine.")
	time.Sleep(5 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)

	go hello(done)
	<-done
	fmt.Println("main function")
}
