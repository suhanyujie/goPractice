package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l1 list.List
	//var e1 = &list.Element {
	//	nil,
	//	nil,
	//	&l1,
	//	12,
	//}
	//_ = l1.PushFront(12);
	//var firstEle = l1.Front();
	for i:=12;i>0;i--{
		l1.PushFront(i)
		fmt.Printf("now List's length is :%d\n", l1.Len())
	}
}


/**
## 一些概念
* 各个类型的零值：只做了声明，但还未做初始化的变量被给予的缺省值，每个类型的零值都会依据该类型的特征而被设定。


## 参考
* container标准库的文档 https://gowalker.org/container/list
* List的其他文档 https://gowalker.org/container/list#New
 */