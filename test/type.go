package main

import (
	"fmt"
	"time"
)

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

	//time
	printTime()
}

func printTime(){
	fmt.Println("=========printTime")
	t := time.Now()
	str := fmt.Sprintf("%v-%04d-%02d-%02d.log", "filename", t.Year(), t.Month(), t.Day())
	fmt.Println(str)
}
