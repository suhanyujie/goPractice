package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	testStr := "hi Samuel"
	md5Obj := md5.New()
	md5Obj.Write([]byte(testStr))
	result := md5Obj.Sum([]byte("1"))
	fmt.Printf("%x\n", result)

	shalInst := sha1.New()
	shalInst.Write([]byte(testStr))
	result = shalInst.Sum([]byte(""))
	fmt.Printf("%x\n", result)
}
