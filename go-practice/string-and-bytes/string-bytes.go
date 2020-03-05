package strings_and_bytes

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func Contains() {
	fmt.Printf("%v\n", strings.Contains("hello world", "gophers"))
	fmt.Printf("%v\n", bytes.Contains([]byte("abc"), []byte("b")))
}

func TrimFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fmt.Printf("[%q]\n", strings.TrimFunc("    23venkat gopher402  ", f))
}

func Replacer() {
	r := strings.NewReplacer("alpha", "A", "theta", "Θ", "delta", "Δ")
	fmt.Printf("%q\n", r.Replace("The alpha differs from the theta which differs from the delta."))
}

func ReplacerForBytes() {
	fmt.Println("Need to implement")
	fmt.Println("and convert string to byte and vice versa")
}

func Join() {
	alphabet := "  alpha beta gamma  "
	fields := strings.Fields(alphabet)
	fmt.Println("Fields %q", fields)
}

func Reader() {
	var a, b, c string
	s := "a b c"
	fmt.Printf("Before scanning: %s, %s, %s\n", a, b, c)
	r := strings.NewReader(s)
	if _, err := fmt.Fscan(r, &a, &b, &c); err != nil {
		fmt.Println("Got error while scanning reader", err)
		return
	}
	fmt.Printf("After scanning: %s, %s, %s\n", a, b, c)
}
