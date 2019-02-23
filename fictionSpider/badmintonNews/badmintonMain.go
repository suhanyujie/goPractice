package main

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"errors"
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
)

type OneNews struct {
	Title string
	Link string
}

// 入口函数
func main() {
	url := "http://www.badmintoncn.com/"
	status,content, _ := GetTopNews(url)
	fmt.Println(status)
	fmt.Println(content)
}

/**
@desc 获取小说的详情
 */
func GetTopNews(url string) (status interface{}, content []OneNews, returnErr error) {
	result, err := GetHttpResponse(url)
	if err != nil {
		return 30021, []OneNews{}, err
	}
	//解析html内容
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(result))
	if err != nil {
		return 30026, []OneNews{}, err
	}
	var topNews []OneNews
	var detailContent,link string
	var isHave bool
	var TmpNews = new(OneNews)
	dom.Find("#slideshow .image li").Each(func(i int, s *goquery.Selection) {
		detailContent,err = s.Find("div").Html()
		link,isHave = s.Find("a").Attr("href")
		TmpNews.Title = detailContent
		if isHave {
			TmpNews.Link = link
		} else {
			TmpNews.Link = ""
		}
		topNews = append(topNews, *TmpNews)
	});

	//dir, _ := os.Getwd()
	//newFileName := dir + "/src/practice/http/testSpider/cache/fiction.inc"
	//将内容写入文件中
	//fileContent := []byte(detailContent)
	//err = file.SsWriteFile(newFileName, fileContent)
	//if err != nil {
	//	return 30036,"", err
	//}

	return 1, topNews, errors.New("任务完成...")
}

/**
根据url请求对应的详细并返回
 */
func GetHttpResponse(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(err.Error() + "-------[30015]")
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_13_6) AppleWebKit/537.36 (KHTML,Like Gecko)Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New(err.Error() + "-------[30021]")
	}
	defer response.Body.Close()
	//fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.New(fmt.Sprintf("该请求的状态码为：%d\n", response.StatusCode))
	}
	utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

	return ioutil.ReadAll(utf8Content)
}


