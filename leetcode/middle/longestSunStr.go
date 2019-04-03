package main

import "fmt"

/**
## 最长子串
* 地址 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/

func main() {
	//test1()
	originStr := "dvdf" //aadvdf  dvdf    bbbbb    helloworld
	longestLen := lengthOfLongestSubstring(originStr)
	fmt.Println(longestLen)
}

func test1() {
	originStr := "dvdf"
	l1 := len(originStr)
	for i := 0; i < l1; i++ {
		fmt.Println(i)
	}
}

type SubStr struct {
	Len uint8
	Con string
}

func lengthOfLongestSubstring(s string) int {
	originStr := s
	resStr := ""
	var (
		b1             []byte
		longestObj     SubStr
		tmpLen         int
		l1             int
		curChar        uint8
		firstRealIndex int
	)
	firstRealIndex = 0
	originStrByte := []byte(originStr)
	l1 = len(originStr)
	for i := 0; i < l1; i++ {
		curChar = originStrByte[i]
		// invalidIndex 表示出现相同字符在原字符串中的位置
		isValid, invalidIndex := InputVali(int32(curChar), resStr, i, &firstRealIndex)
		// 发生冲突后，先假设invalidIndex 就是发生冲突的位置p1
		if !isValid {
			// 一旦不合法，则先打印一次，然后清除bufer b1
			// 不合法时，先找到冲突的字符所在位置，回溯至其位置之后的一个字符，开始新一轮的字符对比
			resStr = string(b1)
			tmpLen = len(resStr)
			if int(longestObj.Len) < tmpLen {
				longestObj.Len = uint8(tmpLen)
				longestObj.Con = resStr
			}
			if invalidIndex == i-1 {
				// curChar此时是int32
				b1 = []byte{}
				b1 = append(b1, uint8(curChar))
				resStr = string(b1)
				// 如果是相邻字符重复，则直接赋值当前字符位置
				firstRealIndex = i
			} else {
				// 此时本应i = invalidIndex + 1，但考虑到下一次循环会递增，所以不用+1
				i = invalidIndex
				b1 = []byte{}
				resStr = ""
				// 如果是非相邻字符重复
				firstRealIndex = invalidIndex + 1
			}
		} else {
			b1 = append(b1, uint8(curChar))
			// 字符的连接
			resStr = string(b1)
		}
	}
	resStr = string(b1)
	tmpLen = len(resStr)
	if int(longestObj.Len) < tmpLen {
		longestObj.Len = uint8(tmpLen)
		longestObj.Con = resStr
	}
	// fmt.Println("最长的字串为：", longestObj.Con, " 其长度为：", longestObj.Len)
	return int(longestObj.Len)
}

func InputVali(oneChar int32, resStr string, realIndex int, firstIndex *int) (bool, int) {
	if len(resStr) == 0 {
		return true, -1
	}
	for _, curChar := range resStr {
		if curChar == oneChar {
			return false, *firstIndex
		}
	}

	return true, -1
}
