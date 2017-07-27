package main

import (
	"fmt"
	"reflect"
	"github.com/songcf/hellogo/utils"
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
	reflectField()
	reflectCallNotExportObjFunc()
}

func reflectField() {
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

func reflectCallNotExportObjFunc() {
	obj := utils.GetObj()
	v := reflect.ValueOf(obj).FieldByName("Integ")
	fmt.Println("v:", v)

	// error
	reflect.ValueOf(obj).MethodByName("PrintInt").Call([]reflect.Value{})
	//ok   -> struct里面全是interface才会这样。。。反射必须拿到inteface字段后调用接口 ;;;why?
	reflect.ValueOf(obj).FieldByName("Integ").MethodByName("PrintInt").Call([]reflect.Value{})
}

