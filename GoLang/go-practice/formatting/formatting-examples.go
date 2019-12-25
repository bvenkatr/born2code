package test

import (
	fmt "fmt"
	ioutil "io/ioutil"
)

var (
	i = 20
	f = 2.903
	s = "Clear is better than clever"
	b = true
	p = struct{ x, y int64 }{256, 314}
)

func OutputToWriter() (string, error) {
	file, err := ioutil.TempFile("", "example")
	if err != nil {
		return "", err
	}
	defer file.Close()

	fmt.Fprintf(file, "i = %v\nf = %v\ns = %s\nb = %v\np = %v\nbi = %v\nui = %v\n", i, f, s, b, p, bi, ui)
	return file.Name(), nil
}

var bi int64 = -83948723884
var ui uint64 = 4590349530945

func OutputToStdout() {
	fmt.Println(i, f, s, b, p, bi, ui)
	fmt.Printf("i = %v\nf = %v\ns = %s\nb = %v\np = %v\nbi = %v\nui = %v\n", i, f, s, b, p, bi, ui)
	fmt.Println()
	fmt.Printf("i = %#v\nf = %#v\ns = %#v\nb = %#v\np = %#v\nbi = %#v\nui = %#v\n", i, f, s, b, p, bi, ui)
}
