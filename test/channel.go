package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fun1(ch)
	ch <- 1
	fmt.Println(ch)
	time.Sleep(time.Second * 2)
	close(ch)
	//ch <- 2  //FIXME panic: send on closed channel

	fmt.Println(ch)
	ch = nil
	fmt.Println(ch)

	//double close

}


func fun1(ch chan int) {
	for {
		i := <- ch
		fmt.Printf("i:%v\n", i)
	}
}



