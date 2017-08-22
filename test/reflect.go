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
type user2 struct {
	name string `json:"nm"`
	age int `json:"ag"`
	child []int  `json:"ch"`
	TAuth
}
type TAuth struct {
	UID      int64  `json:"uid"`
	Ticket   string `json:"ticket"`
	UnitID   int64  `json:"unit_id"`
	UAccount string `json:"uid_str"`
}

func (u *user) getName() string {
	return u.name
}

func main() {
	//reflectField()
	//reflectCallNotExportObjFunc()
	//reflectInterface()
	child()
}

func child() {
	a := TAuth{
		UID:1,
		Ticket:"ticket",
		UnitID:11,
		UAccount:"account",
	}
	var u = user2{name:"songcf", age:25, TAuth:a}
	var uu interface{} = &u
	rv := reflect.ValueOf(uu).Elem()
	childV := rv.FieldByName("TAuth")

	rt := reflect.TypeOf(uu).Elem()
	stt, _ := rt.FieldByName("TAuth")
	fmt.Println(stt.Type.Name())

	fmt.Println(childV.FieldByName("Ticket").String())

}

func reflectField() {
	var u = user{name:"songcf", age:25}
	var v = reflect.ValueOf(u)

	if v.FieldByName("TAuth").Kind() == reflect.Invalid {
		var fieldNum = v.NumField()
		for i:=0; i<fieldNum; i++ {
			fmt.Println(v.Type().Field(i).Name, "(",
				v.Type().Field(i).Type, ") : ",
				v.Field(i), ", ",
				v.Type().Field(i).Tag)
		}
	}
}

func reflectInterface() {
	a := TAuth{
		UID:1,
		Ticket:"ticket",
		UnitID:11,
		UAccount:"account",
	}
	var u = user2{name:"songcf", age:25, TAuth:a}
	var uu interface{} = &u
	var rv = reflect.ValueOf(uu).Elem()

	var te interface{} = a.Ticket
	fmt.Println("......kind1: ",reflect.ValueOf(uu).Kind())
	fmt.Println("......kind2: ", reflect.ValueOf(te).Kind())
	fmt.Println("......kind3: ", reflect.ValueOf(a.Ticket).Kind())
	fmt.Println("......kind4: ", reflect.ValueOf(a).Kind())
	fmt.Println("......kind5: ", reflect.ValueOf(a).Kind())

	if rv.Kind() != reflect.Struct {
		panic("reflectInterface error!!!!")
	}
	auth := rv.FieldByName("TAuth")
	if auth.Kind() != reflect.Invalid {
		str := ""
		for j:=0; j<auth.NumField(); j++ {
			str += fmt.Sprintf("%v:%v\n", auth.Type().Field(j).Name, auth.Field(j))
		}
		fmt.Println("str: ", str)
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

