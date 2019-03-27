package main

import (
	"fmt"
	"reflect"
)

func main() {
	//ChangeValue()
	StructReflect()
}

func StructReflect()  {
	type T struct {
		Name string
		Age int
	}
	t := T{"samuel",12}
	s := reflect.ValueOf(&t).Elem()
	type1 := s.Type()
	for i:=0;i<s.NumField() ;i++  {
		fName := s.Field(i)
		// 如果结构体属性不是public，则不能通过反射获取其值！
		fmt.Printf("%d:%s %s = %+v \n",
			i, type1.Field(i).Name,fName.Type(),fName.Interface())
	}
}

func ChangeValue()  {
	var x float64 = 3.14
	p := reflect.ValueOf(&x)
	v := p.Elem()
	fmt.Println(v.CanSet())
	v.SetFloat(1.111)
	fmt.Println(x)
}

func showValueAndType()  {
	var x float64 = 3.14
	v := reflect.ValueOf(x)

	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v)
}