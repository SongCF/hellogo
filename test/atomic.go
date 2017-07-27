package main

import (
	"sync/atomic"
	"os"
	"time"
	"fmt"
)




func main() {
	var v atomic.Value

	f1, err := os.OpenFile("test1.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("open file failed!")
	}
	v.Store(f1)
	f1.WriteString("test string")
	time.Sleep(time.Second * 2)

	f1s, e := f1.Stat()
	fmt.Printf("f1 state: %v,%v\n", f1s, e)

	f2, err := os.OpenFile("test1.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("open file failed!")
	}
	v.Store(f2)
	f2.WriteString("test string")
	time.Sleep(time.Second * 2)

	f1s, e = f1.Stat()
	fmt.Printf("f1 state: %v,%v\n", f1s, e)
	f1.Close()

	time.Sleep(time.Second * 2)
}