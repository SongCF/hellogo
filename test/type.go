package main

import "fmt"

type ttt struct {
	i int
	str string
}

func main() {
	i := 10
	pi := &i
	t, ok := interface{}(pi).(*ttt)
	if ok {
		fmt.Println("true")
		fmt.Println(t)
	} else {
		fmt.Println("false")
	}
}
