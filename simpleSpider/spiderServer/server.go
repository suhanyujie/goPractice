package spiderServer

import (
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type InputData struct {
	Text string
	Rule string
}

type SpiderServer interface {
	HandleReceive(text, rule string) (result string, err error)
}

type NovelServer struct {
	DataChan *chan InputData
}

func (server *NovelServer) HandleReceive(text, rule string, textRule []string) {
	textReader := strings.NewReader(text)
	doc, err := goquery.NewDocumentFromReader(textReader)
	if err != nil {
		log.Fatal(err)
	}
	ruleByte, err := base64.StdEncoding.DecodeString(rule)
	if err != nil {
		log.Println("spider rule decode 异常！")
		return
	}
	rule = string(ruleByte)
	doc.Find(rule).Each(func(i int, selection *goquery.Selection) {

	})
}

func NewNovelServer() *NovelServer {
	ch1 := make(chan InputData, 200)
	return &NovelServer{&ch1}
}
