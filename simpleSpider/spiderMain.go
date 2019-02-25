package main

import (
	"practice/simpleSpider/spiderServer"
	"net"
	"log"
	"encoding/base64"
	"bufio"
)

func main() {
	//启动一个tcp服务器，接收请求的数据
	//将数据传递给server
	server := spiderServer.NewNovelServer()
	ln,err := net.Listen("tcp","0.0.0.0:8001")
	if err != nil {
		log.Fatal("tcp server Listen error:", err)
	}
	defer ln.Close()
	for {
		conn,err := ln.Accept()
		if err != nil {
			log.Fatal("tcp server Accept error:", err)
		}

		go Handle(conn, *server)
	}
}

func Handle(conn net.Conn,server spiderServer.NovelServer) {
	var (
		//bNum int
		err error
		dataStr string
	)
	defer (conn).Close()
	//var data = make([]byte, 1024)
	var bufReader = bufio.NewReader(conn)
	for {
		dataStr,err = bufReader.ReadString('\n')
		//bNum,err = (conn).Read(data)
		if err != nil {
			log.Fatal("tcp server Accept error:", err)
		}
		//if bNum < 1 {
		//	fmt.Println("request data is null.")
		//	return
		//}
	}
	//解析数据，放入spider server处理
	server.HandleReceive(dataStr, getAreaRule(), getItemRule())
}

func getAreaRule() string  {
	areaRule := "#list"
	areaRule = base64.StdEncoding.EncodeToString([]byte(areaRule))
	return areaRule
}

func getItemRule() []string {
	var ruleArr = []string{
		"dd a",
	}
	for index, con := range ruleArr {
		newRule := base64.StdEncoding.EncodeToString([]byte(con))
		ruleArr[index] = string(newRule)
	}
	return ruleArr
}

//var chTask = make(chan beegoOrm.FictionOneOfList)

// 入口函数
//func main() {
//	var wg sync.WaitGroup
//	listUrl := "https://www.biduo.cc/biquge/17_17308"
//	//url := "https://www.biduo.cc/biquge/17_17308/c8698877.html"
//	//status,err := testSpider.GetDetail(url)
//	testSpider.GetBaseUrl(listUrl)
//	go func(wg sync.WaitGroup) {
//		wg.Add(1)
//		status, _, err := testSpider.GetList(chTask, listUrl)
//		if status != nil {
//			log.Fatal(err)
//		}
//		logs.Info(err)
//		wg.Done()
//	}(wg)
//	// 将每一个任务数据放入数据库，
//	// 然后放入任务队列中，获取具体的详情页内容
//	go testSpider.DealTask(chTask,wg)
//
//	wg.Wait()
//}
