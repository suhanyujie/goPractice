package spiderServer

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type InputData struct {
	Text string
	Rule string
}

type LinkData struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

type SpiderServer interface {
	HandleReceive(text, rule string) (result string, err error)
}

type NovelServer struct {
	DataChan chan []LinkData
}

func (server *NovelServer) HandleReceive(text, rule string, textRule map[string]string) {
	var (
		oneRule     string
		aSelection  *goquery.Selection
		needLink    bool
		linkData    = &LinkData{"", ""}
		linkDataArr = make([]LinkData, 0)
	)
	text = strings.TrimSpace(text)
	textReader := strings.NewReader(text)
	doc, err := goquery.NewDocumentFromReader(textReader)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("=================解析结果如下：=========================================")
	doc.Find(rule).Each(func(i int, selection *goquery.Selection) {
		for _, oneRule = range textRule {
			//itemRule, err := base64.StdEncoding.DecodeString(oneRule)
			if err != nil {
				log.Println("spider itemRule decode 异常:", err)
				return
			}
			aSelection = selection.Find(oneRule)
			//判断是否a链接
			needLink = strings.HasSuffix(oneRule, "a")
			aSelection.Each(func(i int, s2 *goquery.Selection) {
				linkData.Text = s2.Text()
				if needLink {
					linkData.Link, _ = s2.Attr("href")
				}
				linkDataArr = append(linkDataArr, *linkData)

				//fmt.Println(linkData.Text)
			})
		}
	})
	server.DataChan <- linkDataArr
}

func NewNovelServer() *NovelServer {
	ch1 := make(chan []LinkData, 10)
	return &NovelServer{ch1}
}
