package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if r := recover();r!=nil  {
			log.Printf("Runtime error caught:%v\n", r)
		}
	}()

	function1()

	fmt.Println("hello world...\n")
}

func function1() {
	panic("manual panic...")
}



/**
## recover
* recover()函数用于终止错误处理流程。一般情况下，recover()应该在使用一个defer关键字的函数中执行以有效地截取错误处理流程。
* 如果没有在发生异常的goroutine中明确调用恢复过程（也就是使用recover），会导致该goroutine所属的进程打印异常信息后直接退出！






 */