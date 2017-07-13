package main

import (
	"math/rand"
	"fmt"
	"time"
)


func adjust(list []int, parent int, length int, cpN int) int {
	lchild := 2*parent + 1
	rchild := 2*parent + 2
	max := parent

	cpN++
	if lchild < length {
		if list[max] < list[lchild] {
			max = lchild
			cpN++
		}
		if rchild < length && list[max] < list[rchild] {
			max = rchild
			cpN++
		}
		cpN++
		if max != parent {
			temp := list[parent]
			list[parent] = list[max]
			list[max] = temp
			cpN = adjust(list, max, length, cpN)
		}
	}
	return cpN
}


func main() {
	var compNum = 0
	var count int = 100000
	var top int = 100

	//random array
	rand.Seed(int64(time.Now().Second()))
	a1 := rand.Perm(count*2)
	arr := a1[:count]

	// init heap
	for i := count / 2; i >= 0; i-- {
		compNum = adjust(arr, i, len(arr)-1, compNum)
	}
	fmt.Println("init count: ", compNum)
	//return

	for i := count-1; i > count-1-top; i-- {
		tmp := arr[0]
		arr[0] = arr[i]
		arr[i] = tmp
		fmt.Println(count-i, ": ", arr[i])
		compNum = adjust(arr, 0, i, compNum)
	}
	//fmt.Println(arr)
	fmt.Println("compare num: ", compNum)
}