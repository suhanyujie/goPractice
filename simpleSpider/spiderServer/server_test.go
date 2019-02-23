package spiderServer

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNovelServer(t *testing.T) {
	html, err := getHtmlContent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)
	//server := NewNovelServer()
	//server.HandleReceive("", "", []string{})
}

func getHtmlContent() ([]byte, error) {
	path := "D:\\vmShare\\go\\src\\practice\\simpleSpider\\spiderServer\\test.txt"
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var html []byte
	n, err := file.Read(html)
	fmt.Println("the content length is:", n)
	if err != nil {
		return nil, err
	}

	return html, nil
}
