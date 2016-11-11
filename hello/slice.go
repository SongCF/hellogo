package main

import "fmt"

func Slice() {
	s := make([]int, 10)
	fmt.Println(s)
	s[7] = 7
	fmt.Println("len = ", len(s), ", cap = ", cap(s))
	s = s[2:4]
	fmt.Println("len = ", len(s), ", cap = ", cap(s))

	//string
	var str1 string = "abcdef"
	sub_str1 := str1[2:4]
	fmt.Println(sub_str1)
	fmt.Println("")
}