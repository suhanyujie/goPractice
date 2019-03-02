package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

// 运行这个文件的代码时，先到项目跟目录，编译：`go build -o a.exe goLanguageCode/chapter6/hash/hash2.go`
// 再运行：`./a.exe`

func main() {
	testFile := "goLanguageCode/chapter6/hash/test.inc"
	infile, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}
	md5Ins := md5.New()
	io.Copy(md5Ins, infile)
	fmt.Printf("%x %s\n", md5Ins.Sum([]byte("")), testFile)
}
