package utils

import "fmt"


type Obj interface {
	Print() error
}
type Integ interface {
	PrintInt() error
}

type notExportObj struct {
	Integ
	Obj
}
func (o *notExportObj) PrintInt() error {
	return o.Integ.PrintInt()
}


type Str struct {
	str string
}
func (obj *Str) Print() error {
	fmt.Println("Str: ", obj.str)
	return nil
}

type Int struct {
	i int
}
func (obj *Int) PrintInt() error {
	fmt.Println("Int: ", obj.i)
	return nil
}

//用来测试反射
func GetObj() Obj {
	obj := &Str{"string"}
	it := &Int{1}
	return notExportObj{it, obj}
}