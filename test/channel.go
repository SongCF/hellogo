package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan []byte)
	go fun2(ch1)
	data := []byte("test")
	ch1 <- data   //等到func2取走数据后才会执行下一步
	fmt.Println("ch <- 1")
	//如果没有设置容量，或者容量设置为0, 说明Channel没有缓存，
	//只有sender和receiver都准备好了后它们的通讯(communication)才会发生(Blocking)。
	// 如果设置了缓存，就有可能不发生阻塞， 只有buffer满了后 send才会阻塞， 而只有缓存空了后receive才会阻塞。
	ch1 <- data
	fmt.Println("ch <- 2")
	close(ch1)


	fmt.Println("=============")
	ch2 := make(chan int)
	go fun1(ch2)
	ch2 <- 1
	fmt.Println(ch2)
	time.Sleep(time.Second * 2)
	close(ch2)
	//ch <- 2  //panic: send on closed channel

	fmt.Println(ch2)
	ch2 = nil
	fmt.Println(ch2)
}


func fun1(ch chan int) {
	for {
		i := <- ch
		fmt.Printf("i:%v\n", i)
	}
}

func fun2(ch chan []byte) {
	for {
		fmt.Println("fun2:", <-ch)
		time.Sleep(time.Second * 10)
	}
}



