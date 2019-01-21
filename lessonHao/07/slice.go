package main

import "fmt"

func main() {
	//指明切片长度
	s1 := make([]int,5);
	fmt.Printf("s1 value's len is :%d\n", len(s1));
	fmt.Printf("s1 value's cap is :%d\n", cap(s1));
	fmt.Printf("s1 value is :%d\n", s1);
	//指明切片长度 和 切片容量
	s2 := make([]int, 5, 8)
	fmt.Printf("s2 value's len is :%d\n", len(s2));
	fmt.Printf("s2 value's cap is :%d\n", cap(s2));
	fmt.Printf("s2 value is :%d\n", s2);




}
/**
## 07 数组和切片
* 理解字面量、值字面量、类型字面量
* 数组是切片的底层数组，切片是数组的某个连续部分的引用
* 数组属于值类型，而切片属于引用类型
* 使用make初始化切片时，如果不指明切片容量，则其容量和长度是一致的
* 切片的容量实际代表了其底层数组的长度，这只是针对make函数初始化的切片，不适合"切片表达式"生成的切片
* 切片可以看做数组封装后的窗口，这个窗口只能在数组上向右扩张，不能向左，所以有时候，切片的容量是小于其底层数组的长度
* 切片的容量总会是在切片长度和底层数组长度之间的某一个值
* 切片的扩容，不会改变原切片，也会生成一个新的底层数组，从而重新生成一个新的切片
* 可以简单地认为新切片的容量（以下简称新容量）将会是原切片容量（以下简称原容量）的 2 倍。
* 当原切片的长度（以下简称原长度）大于或等于1024时，Go 语言将会以原容量的1.25倍作为新容量的基准（以下新容量基准）
* 在对切片进行追加内容时，只要不超过其容量，就不会引发扩容
* 问：如何使用扩容的思想，对切片进行缩容？

 */