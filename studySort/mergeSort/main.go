package main

import "fmt"

/**
## 归并排序：
* 归并排序（MERGE-SORT）是利用归并的思想实现的排序方法，该算法采用经典的分治（divide-and-conquer）策略（分治法将问题分(divide)成一些小的问题然后递归求解，而治(conquer)的阶段则将分的阶段得到的各答案"修补"在一起，即分而治之)。



 */
// 入口函数
func main() {
	var dataArr = []int{12,7,52,58,9,21,53,68,49}
	RecursiveSort(&dataArr)
	fmt.Println(dataArr)
}

//递归方式实现的归并排序
func RecursiveSort(dataArr *[]int) []int {
	len1 := len(*dataArr)
	//先差UN个建一个长度等于原数组的临时数组
	var (
		tmpArr [len1]int
	)



	return nil
}

func Sort() {

}

