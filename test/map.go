package main

import "fmt"

type t struct {
	m map[int]int
}

var gm = make(map[int]*t)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	//b := m["b"] == nil   //panic
	//fmt.Printf("%v\n", b)

	fmt.Println("=======range")
	for k,v := range m {
		fmt.Printf("K:%v,v:%v\n", k,v)
	}
	for k := range m {
		fmt.Printf("k:%v\n", k)
	}
	fmt.Println("==============")

	fmt.Println("=======struct init")
	gm[1] = &t{}
	gm[1].m = make(map[int]int) //necessary
	gm[1].m[1] = 1
	fmt.Println(gm[1].m[1])
	fmt.Println("==============")
}