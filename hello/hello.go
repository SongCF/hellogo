package main

import (
	"fmt"
	"github.com/songcf/hellogo/utils"
)


var x, y = 1, 2

func main() {
	fmt.Println("Hello, Golang!")

	// global var
	test1()
	fmt.Println("package global val: x = ", x, ", y = ", y)

	ConstBit()
	Defer()
	Format()
	Goruntime()
	Channel()
	Interface()
	Slice()


	fmt.Println("reverse Hello:", utils.ReverseString("Hello"))
	fmt.Println("======华丽的分割线======")
}

func test1() {
	x = 10
}



