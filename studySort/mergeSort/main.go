package main

import "fmt"

/**
归并排序：



## 参考资料
* https://juejin.im/post/5ab4c7566fb9a028cb2d9126
 */

// 入口函数
func main() {
	var arr = []int{5, 2, 3,12,956,45,453,19,53};
	MergeSort(&arr, 0, 2)
}

/*
归并排序，使用分治思想


*/
func MergeSort(arr *[]int, L, R int) {
	if L == R {
		return
	}
	m := (L + R) / 2
	MergeSort(arr, L, m)
	MergeSort(arr, m+1, R)
	//排好序后 进行合并
	Merge(arr, L, m+1, R)
}

//数据合并
func Merge(arr *[]int, L, M, R int) {
	var resArr = []int{}
	for i := 0; i < M-L || i < R-M; i++ {
		if (*arr)[L+i] < (*arr)[M+i] {
			fmt.Println((*arr)[L+i])
			resArr = append(resArr,(*arr)[L+i])
		} else {
			resArr = append(resArr,(*arr)[M+i])
		}
	}
	fmt.Println(resArr)
}
