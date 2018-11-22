package main

import "fmt"

func swap(a []int, i, j int) {
	var tmp = a[i]
	a[i] = a[j]
	a[j] = tmp
}
func fnFindIdxNum(a []int, s, e, idx int) int {
	var p, l, r = s, s + 1, e
	var tmp = a[p]
	for l <= r {
		if a[l] < tmp {
			a[l-1] = a[l]
			l++
		} else {
			swap(a, l, r)
			r--
		}
	}
	a[l-1] = tmp
	p = l - 1
	//fmt.Println(a, p, idx)
	if idx == p {
		return a[idx]
	} else if p < idx {
		return fnFindIdxNum(a, p+1, e, idx)
	} else {
		return fnFindIdxNum(a, s, p-1, idx)
	}
}

func main() {
	var fnMidNum = func(a []int) float32 {
		if len(a) == 0 {
			return 0
		}
		if len(a)%2 == 0 {
			var m1 = fnFindIdxNum(a, 0, len(a)-1, len(a)/2-1)
			var m2 = fnFindIdxNum(a, 0, len(a)-1, len(a)/2)
			return float32(m1+m2) / 2.0
		} else {
			return float32(fnFindIdxNum(a, 0, len(a)-1, len(a)/2))
		}
	}

	//
	var a = []int{1, 2, 3, 4, 5, 6, 100, 20, 0}
	fmt.Println(a, " => ", fnMidNum(a))
	a = append(a, 10)
	fmt.Println(a, " => ", fnMidNum(a))
}
