package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go p_c(ch)
	for {
		select {
		case i, ok := <- ch:
			if ok {
				fmt.Printf("i:%d\n", i)
				close(ch)
			}
		default:
			//fmt.Println("default")
		}
	}
}

func p_c(ch chan int) {
	for i := 1; ; i++ {
		fmt.Println("======")
		panic_recover(ch, i)
		time.Sleep(time.Second * 3)
	}
}

func panic_recover(ch chan int, i int) {
	defer func() {
		fmt.Println("defer...")
		if err := recover(); err != nil {
			fmt.Println("recover error: ", err)
		}
		fmt.Println("defer...end")
	}()

	ch <- i
}