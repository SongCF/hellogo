package main

import "fmt"


type Int int
func (i Int) Abs() int {
	if i < 0 {
		return int(-i)
	} else {
		return int(i)
	}
}

type Abser interface {
	Abs() int
}

func Interface() {
	// type func
	var tf Int = -10
	fmt.Println(tf.Abs())
	fmt.Println("")

	//interface
	var abser Abser
	var absera Int = -1
	abser = absera
	fmt.Println("interface, ", abser.Abs())
	fmt.Println("")
}