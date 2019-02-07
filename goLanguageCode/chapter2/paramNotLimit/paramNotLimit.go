package main

import "fmt"

func main() {
	func1(45,2,3);

}

func func1(param ...int) {
	for _,arg := range param {
		fmt.Println(arg)
	}
}

/**
## 不定参数
* 形如`...type`格式的类型只能作为函数的参数类型存在




 */
