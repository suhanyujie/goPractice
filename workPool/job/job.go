package job

import (
	"hash/fnv"
	"time"
	"fmt"
	"math/rand"
)

/**
学完此文后，需要学：
	Runes。通过学习次博文项目，发现对Runes的理解还是非常少！

 */
var letterRunes  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//模拟任意类型的可以同时发生job
func DoWork(word string,id int) {
	h:= fnv.New32a()
	h.Write([]byte(word))
	time.Sleep(time.Second)
	// os.Getenv("DEBUG") == "true"
	if true {
		fmt.Printf("work [%d]-created hash [%d] from word [%s]\n", id,h.Sum32(),word)
	}
}

/**
创建工作列表
RandStringRunes(8) 模拟的是一些工作，比如http请求，接口请求 mysql查询等等

 */
func CreateJobs(amount int) []string {
	var jobs []string
	for i:=0;i<amount ;i++  {
		jobs = append(jobs, RandStringRunes(8))
	}
	return jobs
}

//创建随机的字符串
func RandStringRunes(n int)string {
	b := make([]rune, n)
	for i:=range b{
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
