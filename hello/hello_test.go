package main

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {
	fmt.Println("!!!test begin!!!")
	main()
	fmt.Println("!!!test end!!!")
	//t.Errorf("test error!")
}