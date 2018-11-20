package main

import "fmt"

func swap(a []int, i, j int) {
	var tmp = a[i]
	a[i] = a[j]
	a[j] = tmp
}

func QSort(a []int, s, e int) {
	var i = s + 1
	var j = e
	var tmp = a[s]
	for i <= j {
		if a[i] < tmp {
			a[i-1] = a[i]
			i++
		} else {
			swap(a, i, j)
			j--
		}
	}
	a[i-1] = tmp
	if s < i-2 {
		QSort(a, s, i-2)
	}
	if i < e {
		QSort(a, i, e)
	}
}

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 100, 20, 0}
	QSort(a, 0, len(a)-1)
	fmt.Println(a)
}
