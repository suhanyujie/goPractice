package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	paramReader := strings.NewReader("a=123")
	req, err := http.NewRequest("GET", "http://www.baidu.com", paramReader)
	if err != nil {
		fmt.Println(err)
	}
	c1 := &http.Cookie{
		"name",
		"samuel",
		"/",
		nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil,
	}
	req.AddCookie(c1)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
}
