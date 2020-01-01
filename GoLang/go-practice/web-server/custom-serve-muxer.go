package web_server

import (
	"fmt"
	"log"
	"net/http"
)

// muxer is a custom request multiplexer
type muxer struct{}

func (m *muxer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/greet":
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello Gopher!")
		}(w, r)
	default:
		http.NotFound(w, r)
	}
}

func CustomServeMux() {
	m := &muxer{}
	log.Println("Starting the server...")
	log.Fatal(http.ListenAndServe("0.0.0.0:9999", m))
}