package main

import (
	"fmt"
	"practice/goLanguageCode/chapter5/dial/common"
)

func Hello() []int {
	return Arr()
}

func Arr() []int {
	return []int{12, 9, 67, 53}
}

func main() {
	fun1 := &Hello()[0]

	fmt.Println(fun1)
	common.LookupHost()

	ip := "127.0.0.1"
	res := common.CheckIsIp(ip)
	fmt.Println(res)
}
