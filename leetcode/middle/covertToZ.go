package main

import "log"

func main() {
	str := "helloworld"
	convert(str, 3)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	b1 := []byte(s)
	b1Len := len(b1)
	arr := make([][]byte, 10, 10)
	curRow := 0
	goingDown := false
	for i := 0; i < b1Len; i++ {
		arr[curRow] = append(arr[curRow], b1[i])
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}

	log.Println(arr)

	return ""
}
