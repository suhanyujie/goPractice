package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	resByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resByte))
}
