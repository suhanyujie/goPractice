package main

import "fmt"

// 入口函数
func main() {



     fmt.Println("end---")
}

//遍历字符串
func TraverseStr() {
	str1 := "nishishei suhanyu"
	var len1 = len(str1)
	for i:=0;i<len1;i++{
		fmt.Printf("index:%d,val:%c\n", i, str1[i])
	}
}
