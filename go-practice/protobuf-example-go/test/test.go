package test

import (
	fmt "fmt"
)

var (
	i = 20
	f = 2.903
	s = "Clear is better than clever"
	b = true
	p = struct{ x, y int64 }{256, 314}
)

var bi int64 = -83948723884
var ui uint64 = 4590349530945

func OutputToStdout() {
	fmt.Println(i, f, s, b, p, bi, ui)
	fmt.Printf("i = %v\nf = %v\ns = %s\nb = %v\np = %v\nbi = %v\nui = %v\n", i, f, s, b, p, bi, ui)
	fmt.Println()
	fmt.Printf("i = %#v\nf = %#v\ns = %#v\nb = %#v\np = %#v\nbi = %#v\nui = %#v\n", i, f, s, b, p, bi, ui)
}
