package main

import (
	"log"
)

/**
最长回文子串 https://leetcode-cn.com/problems/longest-palindromic-substring/solution/
*/
func main() {
	s1 := "aabcba"
	lStr := longestPalindrome(s1)
	log.Println(lStr)
	//TestLongestPalindrome()
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	l1 := len(s)
	start := 0
	end := 0
	compareLen := 0
	maxLen := 0
	len1 := 0
	len2 := 0
	for i := 0; i < l1; i++ {
		len1 = expandAroundCenter(s, i, i)
		len2 = expandAroundCenter(s, i, i+1)
		if len1 > len2 {
			compareLen = len1
		} else {
			compareLen = len2
		}
		if compareLen >= maxLen {
			start = i - (compareLen-1)/2
			end = i + compareLen/2
			maxLen = compareLen
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, l, r int) int {
	L := l
	R := r
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}

	return R - L - 1
}

func TestLongestPalindrome() {
	var exmStr = ""
	resStr := longestPalindrome(exmStr)
	log.Println(resStr)
}
