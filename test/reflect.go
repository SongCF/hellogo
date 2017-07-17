package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string `json:"nm"`
	age int `json:"ag"`
	child []int  `json:"ch"`
}

func (u *user) getName() string {
	return u.name
}

func main() {
	var u = user{name:"songcf", age:25}
	var v = reflect.ValueOf(u)
	var fieldNum = v.NumField()
	for i:=0; i<fieldNum; i++ {
		fmt.Println(v.Type().Field(i).Name, "(",
			v.Type().Field(i).Type, ") : ",
			v.Field(i), ", ",
			v.Type().Field(i).Tag)
	}
}
