package main

import (
	"fmt"
	"time"
)

func say(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second / 5)
		fmt.Println(str)
	}
}

func Goruntime() {
	go say("world")
	//fmt.Println("hello")
	say("hello")
}

func Channel() {
	// channel
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
