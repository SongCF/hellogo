package main

import "fmt"

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
}