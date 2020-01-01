package error_examples

import (
	"errors"
	"log"
)


var (
	errBadThing = errors.New("something bad")
	errWorseThing = errors.New("worse thing")
)
func ErrorExamples() {

	//////////////////////////////////////////////////
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Printf("boom() failed %v\n", err)
	//	}
	//}()
	//
	//log.Println("hello gophers")
	//boom()
	/////////////////////////

	/////////////////////////
	err := doSomethingBad()
	//err = doSomethingWorse()

	if err != nil {
		switch err {
		case errBadThing:
			log.Printf("Uh oh: %s\n", err)
		case errWorseThing:
			log.Printf("Abandon ship!\n")
		}
	}
	/////////////////////////
}

func boom() {
	panic("BOOM!")
}

func doSomethingBad() error {
	return errBadThing
}

func doSomethingWorse() error {
	return errWorseThing
}