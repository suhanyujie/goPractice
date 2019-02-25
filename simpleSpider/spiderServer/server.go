package spiderServer

import (
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"fmt"
)

type InputData struct {
	Text string
	Rule string
}

type LinkData struct {
	Text,Link string
}

type SpiderServer interface {
	HandleReceive(text, rule string) (result string, err error)
}

type NovelServer struct {
	DataChan chan LinkData
}

func (server *NovelServer) HandleReceive(text, rule string, textRule []string) {
	fmt.Println(text)
	text = strings.TrimSpace(text)
	textReader := strings.NewReader(text)
	doc, err := goquery.NewDocumentFromReader(textReader)
	if err != nil {
		log.Fatal(err)
	}
	ruleByte, err := base64.StdEncoding.DecodeString(rule)
	if err != nil {
		log.Println("spider rule decode 异常:", err)
		return
	}
	rule = string(ruleByte)
	doc.Find(rule).Each(func(i int, selection *goquery.Selection) {
		for _, oneRule := range textRule {
			itemRule, err := base64.StdEncoding.DecodeString(oneRule)
			if err != nil {
				log.Println("spider itemRule decode 异常:", err)
				return
			}
			itemRuleStr := string(itemRule)
			aSelection := selection.Find(itemRuleStr)
			//判断是否a链接
			needLink := strings.HasSuffix(itemRuleStr,"a")
			var (
				linkData = &LinkData{"",""}
			)
			aSelection.Each(func(i int, s2 *goquery.Selection) {
				linkData.Text = s2.Text()
				if needLink {
					linkData.Link,_ = s2.Attr("href")
				}
				server.DataChan <- *linkData
				fmt.Println(linkData)
			})
		}
	})
}

func NewNovelServer() *NovelServer {
	ch1 := make(chan LinkData, 200)
	return &NovelServer{ch1}
}
