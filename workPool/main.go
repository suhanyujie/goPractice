package main

import (
	"practice/workPool/job"
	"fmt"
)

// 入口函数
func main() {
	var s1 = job.CreateJobs(12)
	fmt.Println(s1[0])
	job.DoWork(s1[0],0)
}
