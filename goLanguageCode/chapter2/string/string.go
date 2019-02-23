package main

import "fmt"

func main() {
	str1 := "hello 世界。。。"
	for i,c :=  range str1{
		fmt.Printf("index:%d content:%V\n", i,c)
	}

}
