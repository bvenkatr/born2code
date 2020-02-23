package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	done := make(chan int)
	go sendData(done)
	fmt.Println(<-done)
}
