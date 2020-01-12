package checsum_and_content_hash

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

//func CheckSum() {
//	h := sha256.New()
//	h.Write([]byte("Hello world!"))
//	fmt.Printf("%x\n", h.Sum(nil))
//}

func CheckSum() {
	f, err := os.Open("./checsum_and_content_hash/test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatalln(err)
	}

	log.Printf("%x\n", h.Sum(nil))
}