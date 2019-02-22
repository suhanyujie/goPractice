package main

import "fmt"

var (
	fun1 = func() {
		fmt.Println("hello")
	}
)

func Hello() {
	fmt.Println("hello")
}

func main() {
	var (
		n1 = []int{12, 90, 87, 64, 94, 19}
		n2 = []int{}
	)

	n2 = n1[0:len(n1)]
	fmt.Println(n2)

	funTmp := &fun1
	funTmp2 := &fun1

	//fmt.Println(&funTmp == &funTmp2)
	fmt.Println(&funTmp)
	fmt.Println(&funTmp2)
}
