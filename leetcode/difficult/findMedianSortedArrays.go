package main

import "fmt"

/**
找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
*/
func main() {
	arr1 := []int{}
	arr2 := []int{3, 4}
	mRes := findMedianSortedArrays(arr1, arr2)
	fmt.Println(mRes)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	totalLen := len1 + len2
	allArr := make([]int, totalLen)
	if len1 == 0 {
		allArr = nums2
	} else if len2 == 0 {
		allArr = nums1
	} else {
		i1 := 0
		i2 := 0
		for i := 0; i < totalLen; i++ {
			if nums1[i1] < nums2[i2] {
				allArr[i] = nums1[i1]
				if i1++; i1 >= len1 {
					i++
					allArr = append(allArr[:i], nums2[i2:len2]...)
					break
				}
			} else {
				allArr[i] = nums2[i2]
				if i2++; i2 >= len2 {
					i++
					allArr = append(allArr[:i], nums1[i1:len1]...)
					break
				}
			}
		}
	}

	if totalLen%2 == 0 {
		return float64(allArr[totalLen/2]+allArr[totalLen/2-1]) / 2
	} else {
		return float64(allArr[(totalLen-1)/2])
	}
}
