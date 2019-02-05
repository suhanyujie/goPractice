package main

import "fmt"

func main() {
	var i = 5;
	switch i {
	case 0:
		fmt.Println(i)
	case 2:
		fmt.Println(i)
	case 4,5,6:
		fmt.Println("4,5,6")
		fallthrough
	default:
		fmt.Println("Default")
	}

}

/**
## 选择语句
* switch左花括号与switch处于同一行
* 条件表达式不限制为常量或者整数
*



 */
