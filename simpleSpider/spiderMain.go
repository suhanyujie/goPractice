package main

import (
	"practice/simpleSpider/spiderServer"
)

func main() {
	server := spiderServer.NewNovelServer()
	server.HandleReceive("", "", []string{})
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
