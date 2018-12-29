package main

import "fmt"

var container = []string{"zero", "one", "two"}

// 入口函数
func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Println(container)
	value, ok := interface{}(container).([]string)
	fmt.Println(ok)
	fmt.Println(value)
	fmt.Println(string(-1))
}


/**
## 知识点
* 正式说明一下，类型断言表达式的语法形式是`x.(T)`，其中的x代表要被判断类型的值。这个值当下的类型必须是接口类型的值
* 在 Go 语言中，interface{}代表空接口，任何类型都是它的实现类型
* 请记住，一对不包裹任何东西的花括号，除了可以代表空的代码块之外，还可以用于表示不包含任何内容的数据结构（或者说数据类型）。
* 所谓类型字面量，就是用来表示数据类型本身的若干个字符。
* 搞清楚别名类型声明与类型再定义，如：

```golang
type MyString2 string
type MyString = string
```

* 两个值潜在类型相同，却属于不同类型，它们之间是可以进行类型转换。但它们的值之间不能进行判等或比较，它们的变量之间也不能赋值



、



### ask 1:
怎样在打印其中元素之前，正确判断变量container的类型...
answer 1:


极客时间版权所有: https://time.geekbang.org/column/article/13601
 */
