package main

import "fmt"

func main() {
	var x, y, w, h float32
	x = 493.999
	y = 57.6
	w = 350
	h = 350
	floor := int(x)
	fmt.Printf("%v, %v, %v, %v\n", x / w, y / h, int(x / w), floor)
	fmt.Println(x)
}