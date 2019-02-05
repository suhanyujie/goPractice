package main

import "fmt"

func main() {
	var i = 120
	for i>0 {
		fmt.Println(i)
		if i == 112 {
			goto LABEL1
		}
		i--
	}
	LABEL1:
		fmt.Println("has goto someone label...")

}

/**
## goto语句
* 它的作用很简单，跳转到本函数内的某个标签


 */
