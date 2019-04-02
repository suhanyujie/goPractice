package main

import "fmt"

/**
## 最长子串
* 地址 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/

type SubStr struct {
	Len uint8
	Con string
}

func main() {
	//test1()
	originStr := "dvdf" // dvdf    bbbbb    helloworld
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

func lengthOfLongestSubstring(s string) int {
	originStr := s
	resStr := ""
	var (
		tmpChar1    uint8
		b1          []byte
		longestObj  SubStr
		tmpLen      int
		l1          int
		curChar     uint8
		originIndex int
		offset      int
		isFirst     int
	)
	isFirst = 1
	originStrByte := []byte(originStr)
	l1 = len(originStr)
	for i := 0; i < l1; i++ {
		curChar = originStrByte[i]
		isValid, invalidIndex := InputVali(int32(curChar), resStr, &isFirst)
		// 发生冲突后，先假设invalidIndex 就是发生冲突的位置p1
		if !isValid {
			if isFirst != 0 {
				offset = i + invalidIndex
			} else {
				offset = 0
			}
			originIndex = i
			// 一旦不合法，则先打印一次，然后清除bufer b1
			// 不合法时，先找到冲突的字符所在位置，回溯至其位置之后的一个字符，开始新一轮的字符对比
			resStr = string(b1)
			tmpLen = len(resStr)
			if int(longestObj.Len) < tmpLen {
				longestObj.Len = uint8(tmpLen)
				longestObj.Con = resStr
			}
			if offset+invalidIndex == originIndex {
				// curChar此时是int32
				tmpChar1 = uint8(curChar)
				b1 = []byte{}
				b1 = append(b1, tmpChar1)
				resStr = string(b1)
			} else {
				// 此时本应i = invalidIndex + 1，但考虑到下一次循环会递增，所以不用+1
				i = invalidIndex
				b1 = []byte{}
				resStr = ""
			}
		} else {
			tmpChar1 = uint8(curChar)
			b1 = append(b1, tmpChar1)
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

func InputVali(oneChar int32, resStr string, isFirst *int) (bool, int) {
	if len(resStr) == 0 {
		return true, -1
	}
	for index, curChar := range resStr {
		if curChar == oneChar {
			*isFirst = 0
			return false, index
		}
	}

	return true, -1
}
