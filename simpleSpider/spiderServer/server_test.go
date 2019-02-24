package spiderServer

import (
	"fmt"
	"log"
	"os"
	"testing"
	"encoding/base64"
	"sync"
)

func TestNovelServer(t *testing.T) {
	var w sync.WaitGroup
	html, err := getHtmlContent()
	if err != nil {
		log.Fatal(err)
	}
	server := NewNovelServer()
	defer close(server.DataChan)
	var ruleArr = []string{
		"dd a",
	}
	for index, con := range ruleArr {
		newRule := base64.StdEncoding.EncodeToString([]byte(con))
		ruleArr[index] = string(newRule)
	}
	areaRule := "#list"
	areaRule = base64.StdEncoding.EncodeToString([]byte(areaRule))
	go func(DataChan <-chan LinkData) {
		w.Add(1)
		var (
			dataOne = &LinkData{}
		)
		for {
			select {
			case *dataOne = <-DataChan:
				fmt.Println(dataOne.Text)
			}
		}
		w.Done()
	}(server.DataChan)
	server.HandleReceive(string(html), areaRule, ruleArr)
	//等待协程执行任务完成
	w.Wait()
}

func getHtmlContent() ([]byte, error) {
	path := "/www/2017/go/src/practice/simpleSpider/spiderServer/test.txt"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var html = make([]byte, 1024000)
	n, err := file.Read(html)
	fmt.Println("the content length is:", n)
	if err != nil {
		return nil, err
	}

	return html, nil
}
