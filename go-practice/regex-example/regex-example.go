package regex_example

import (
	"fmt"
	"regexp"
)

func  Match()  {
	p := "(foo|faux)bar"

	s1 := "fallbar"
	match, err := regexp.MatchString(p, s1)
	fmt.Printf("MatchingString - %v, %v\n", match, err)

	s2 := "foobar"
	match, err = regexp.MatchString(p, s2)
	fmt.Printf("MatchingString - %v, %v\n", match, err)

	s3 := "bar"
	match, err = regexp.MatchString(p, s3)
	fmt.Printf("MatchingString - %v, %v\n", match, err)

	p = "(C)go*"

	b1 := []byte("Cgo is not Go.")
	match, err = regexp.Match(p, b1)
	fmt.Printf("Match - %v, %v\n", match, err)

	b2 := []byte("Where are all my gophers?")
	match, err = regexp.Match(p, b2)
	fmt.Printf("Match - %v, %v\n", match, err)
}

