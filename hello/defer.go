package main

import "fmt"

func Defer() {
	// defer
	defer fmt.Println("defer end")
	for i := 0; i < 5; i++ {
		defer fmt.Println("defer i: ", i)
	}
	fmt.Println("test defer")
}