package main

import "fmt"

func Add(a,b int) (res int) {
	res = a+b
	fmt.Println(res)
	return
}

func main() {
	ch1 := make(chan int)
	for i:=0;i<10 ;i++  {
		go Add(i,i+2)
	}

	ch1 <- 1
}
