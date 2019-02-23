package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
	Age int
	Profile string
}


func main() {
	stu1 := &Stu {
		"suhanyu",
		18,
		"this is a good world~",
	}
	s := reflect.ValueOf(stu1).Elem()
	typeOf := s.Type()
	for i:=0;i< s.NumField();i++ {
		fmt.Printf("%d:%s %s = %v\n", i,typeOf.Field(i).Name)
	}
	fmt.Println(typeOf)
}
