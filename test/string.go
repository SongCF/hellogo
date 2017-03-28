package main

import (
	"fmt"
)

type Str struct {
	s string
	i int
	b []byte
}

func main() {
	// string equal []byte
	{
		fmt.Println("======string equal []byte")
		var str string = "abc"
		b := []byte("AbC")
		r1 := str == string(b)
		r2 := str == "abc"
		//r3 := "AbC" == b  //panic
		fmt.Printf("r1:%v, r2:%v\n", r1, r2)
		fmt.Println("======================")
	}

	// string to []byte
	{
		fmt.Println("======string to []byte")
		b := []byte("AbC")
		var str string = "abc"
		b = []byte(str)
		fmt.Printf("%s\n", b)
		fmt.Println("======================")
	}

	// empty default
	{
		fmt.Println("======empty default")
		pstr := &Str{}
		fmt.Printf("empty %v\n", pstr.s == "")
		fmt.Printf("i: %v, %d\n", pstr.i, pstr.i)
		fmt.Printf("b: %v, %v\n", pstr.b, pstr.b == nil)
		fmt.Println("======================")
	}

	// string add []byte/int
	{
		fmt.Println("======string add []byte/int")
		var str string = "abc"
		b := []byte("AbC")
		xx := str + string(b)//FIXME can not use "str + b"
		fmt.Printf("%v", xx)
		//i := 1
		//xx = str + i  // FIXME not support
		fmt.Printf("%v\n", xx)
		fmt.Println("======================")
	}

	// string array
	{
		fmt.Println("======string array")
		l := []string{}
		l = append(l, "a")
		l = append(l, "b")
		fmt.Printf("%v\n", len(l))
		for m,n := range l {
			fmt.Println("elem:", m, ", n:", n)
		}
		fmt.Println("======================")

	}

}