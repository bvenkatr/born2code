package web_server

import (
	"fmt"
	"log"
	"net/http"
	"../proverbs"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Clinet...\n")
}

type greetHandler struct{}

func (ph *greetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Gopher!")
}

func WebServer() {
	gh := &greetHandler{}
	http.Handle("/greet/", gh)

	ph := proverbs.NewProverbsHandler()
	http.Handle("/proverbs/", ph)

	http.HandleFunc("/", handler)
	log.Println("Starting server....")
	// nil as second argument means `use http's default request multiplexer behind the sceens
	log.Fatal(http.ListenAndServe("0.0.0.0:9999", nil))
}
