package main

import (
	"time"
	"fmt"
)

func main() {
	i := 0
	i = 1
	 go func() {
		 //time.Sleep(time.Second)
		 fmt.Println(i)
	 }()
	i = 2
	time.Sleep(time.Second*2)
}