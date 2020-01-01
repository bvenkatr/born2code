package error_examples

import "log"

func ErrorExamples() {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("boom() failed %v\n", err)
		}
	}()

	log.Println("hello gophers")
	boom()
}

func boom() {
	panic("BOOM!")
}