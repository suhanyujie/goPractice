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
* 单个case中，可以出现多个结果选项；
* 与C语言等规则相反，Go语言不需要用break来明确退出一个case；
* 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；

 */
