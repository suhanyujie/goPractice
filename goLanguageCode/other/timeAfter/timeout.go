package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 = make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			ch1 <- i + 1
		}
	}()
	fmt.Println(time.Now())
	ch2 := time.After(1 * time.Second)
	select {
	case <-ch1:
		fmt.Println("没有超时\n")
	case v2 := <-ch2:
		fmt.Println(v2)
		fmt.Println("超时1\n")
	case <-time.NewTicker(1 * time.Second).C:
		fmt.Println("超时2\n")
	}
	var chan1 = make(chan int, 0)
	chan1 <- 1
}
