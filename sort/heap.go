package main

import "fmt"

//https://blog.csdn.net/her__0_0/article/details/72511047

func HeapSort(a []int) {
	var count = len(a)
	for i := count/2 - 1; i >= 0; i-- {
		adjustHeap(a, i, count)
	}
	for i := count - 1; i >= 0; i-- {
		swap(a, 0, i)
		adjustHeap(a, 0, i)
	}
}

func swap(a []int, i, j int) {
	var tmp = a[i]
	a[i] = a[j]
	a[j] = tmp
}

func adjustHeap(a []int, i, count int) {
	var j = 2*i + 1
	for j < count {
		if j+1 < count && a[j] < a[j+1] {
			j += 1
		}
		if a[i] >= a[j] {
			break
		}
		swap(a, i, j)
		i = j
		j = 2*i + 1
	}
}

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 100, 20, 0}
	HeapSort(a)
	fmt.Println(a)
}
