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
			break
		}
	}
	select {}
}

func p_c(ch chan int) {
	for i := 1; i <= 3; i++ {
		fmt.Println("======")
		i, inter := panic_recover(ch, i)
		fmt.Printf("i:%v, inter:%v\n", i, inter)
		time.Sleep(time.Second * 3)
	}
}

func panic_recover(ch chan int, i int) (int, interface{}) {
	defer func() {
		fmt.Println("defer...")
		if err := recover(); err != nil {
			fmt.Println("recover error: ", err)
		}
		fmt.Println("defer...end")
	}()

	defer fmt.Println("defer 11111")  //测试panic时该defer会不会执行 ：会执行！

	ch <- i // ！！！！！！panic！！！！！！！！！！！
	return 10, ch  //测试panic之后返回值是什么  : 返回该类型的默认值
}