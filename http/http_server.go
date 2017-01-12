package main

import (
	"fmt"
	"net/http"
	"log"
)

type Hello struct {}
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r)
	fmt.Fprint(w, "Hello, Golang!")
}

func main() {
	// http
	println("http test: http://localhost:4000")
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}

