package main

import (
	"fmt"
	"time"
)

func Add(a,b int) (res int) {
	res = a+b
	fmt.Println(res)
	return
}

func Parse(ch1 <-chan int) (res int)  {
	newVal := <- ch1
	res = newVal

	return
}

func main() {
	ch1 := make(chan int)
	//for i:=0;i<10 ;i++  {
	//	go Add(i,i+2)
	//}


	go func() {
		res := Parse(ch1)
		fmt.Println(res)
	}()
	ch1 <- 2

	close(ch1)
	close(ch1)


	time.Sleep(1*time.Second)
}
