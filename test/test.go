package main

import (
	"fmt"
	"time"
	"log"
)


func main() {
	ch := make(chan int, 1000)
	go update(ch)
	for {
		ch <- 1
		time.Sleep(time.Millisecond * 100)
	}
}



func update(ch <- chan int) {
	var timer <- chan time.Time
	var count int
	var beginOneGroup bool
	log.Println("..............enter AliLog update.................")
	for {
		timer = time.After(time.Second)
		count = 0
		beginOneGroup = false
		log.Println("..............once group begin.................")
		for {
			select {
			case <-ch:
				count++
				if count >= 100 {
					//break  //这里的break只能跳出select
					beginOneGroup = true
					fmt.Println("count > 100")
				}
			case <-timer:
				beginOneGroup = true
				fmt.Println("time out")
			}
			if beginOneGroup {
				break
			}
		}
		log.Println(fmt.Sprintf("............ count %d ..............", count))
		if count > 0 {
			log.Println("=======upload count = ", count)
		}
	}
	log.Println("..............out AliLog update.................")
}
