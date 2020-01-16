package io_and_ioutil

type myReader struct {
	content []byte // the stuff we are going to read from
	position int // index of the byte we're up to in our content
}

//func (r *myReader) Read(b []byte) (n int, err error) {
//
//}