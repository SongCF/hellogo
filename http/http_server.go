package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

const addr = "localhost:4000"

type Hello struct {}
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============handle req:", r)
	fmt.Fprint(w, "Hello, Golang!")
}

func main() {
	// http
	println("http test:", addr)
	var h Hello
	err := http.ListenAndServe(addr, h)
	if err != nil {
		log.Fatal(err)
	}
}

