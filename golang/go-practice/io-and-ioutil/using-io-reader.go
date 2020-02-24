package io_and_ioutil

import (
	"io"
	"log"
	"strings"
)

func StringNewReader() {
	reader := strings.NewReader("this is the stuff I'm reading")
	var result []byte
	buf := make([]byte, 4)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		result = append(result, buf[:n]...)
	}
	log.Printf("%v", string(result))

}
