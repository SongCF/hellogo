package main

import "fmt"

var (
	ch = make(chan int)
)

func main() {
	go func() {
		for {
			fmt.Println("<-ch: ", <-ch)
		}
	}()
	ch <- 1
	close(ch)
	ch <- 2   //panic: send on closed channel


}
