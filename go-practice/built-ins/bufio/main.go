package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// An artificial input source
	// const input = "1234 5678 1234567901234567890"
	const input = "1234,5678,1234567901234567890"

	scanner := bufio.NewScanner(strings.NewReader(input))

	// Custom split function
	_ = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	// Split func for comman onComma
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// fmt.Println(data)
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}

		if !atEOF {
			return 0, nil, nil
		}

		return 0, data, bufio.ErrFinalToken
	}
	// scanner.Split(split)
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid Input %s", err)
	}
}
