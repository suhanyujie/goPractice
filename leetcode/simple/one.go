package main

import (
	"fmt"
)

// todo
func main() {

	//findMedianSortedArrays test
	var (
		nums1   = []int{1,3}
		nums2   = []int{2}
		f1 float64
	)
	f1 = findMedianSortedArrays(nums1,nums2)
	fmt.Println(f1)


	//var (
	//	nums   = []int{1, 2, 4, 7, 8, 12, 32, 84, 26}
	//	resArr []int
	//)
	//resArr = twoSum(nums, 88)
	//fmt.Println(resArr)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		totalLen int
		i, j     int
		res float64
	)
	for _, val := range nums2 {
		nums1 = append(nums1, val)
	}
	totalLen = len(nums1)
	for i = 0; i < totalLen; i++ {
		for j = i; j < totalLen; j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
	//偶数
	if (totalLen)%2 == 0 {
		i := totalLen / 2
		res = float64((nums1[i]+nums1[i+1])/2)
	} else {
		//奇数
		index := (totalLen + 1) / 2 -1
		res = float64(nums1[index])
	}

	return res
}

//两数之和
func twoSum(nums []int, target int) []int {
	var (
		i, j   int
		resArr []int
	)
	for i = 0; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				resArr = append(resArr, nums[i])
				resArr = append(resArr, nums[j])
			}
		}
	}
	return resArr
}
