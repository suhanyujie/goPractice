package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func main() {
	responseByteArr := httpClientForUnlock()
	responseObj := new(Response)
	log.Println(string(responseByteArr))
	err := json.Unmarshal(responseByteArr, responseObj)
	if err != nil {
		log.Println(err)
	}
	log.Println(responseObj.Msg)
}

func httpClientForUnlock() []byte {
	paramReader := strings.NewReader("ip=122.97.175.0")
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://pay.magozfqs.com/Admin_Login_f8029ed832b691cd941facca35e68270.html", paramReader)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return res
}

func httpClient1() {
	paramReader := strings.NewReader("a=123")
	req, err := http.NewRequest("GET", "http://www.baidu.com", paramReader)
	if err != nil {
		fmt.Println(err)
	}
	//c1 := &http.Cookie{
	//	"name",
	//	"samuel",
	//	"/",
	//	nil, nil, nil,
	//	nil, nil, nil, nil,
	//	nil, nil,
	//}
	//req.AddCookie(c1)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
}
