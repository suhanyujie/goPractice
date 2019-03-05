package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"practice/telegramApi/apiClient"
	"practice/telegramApi/common"
)

func main() {
	params := map[string]string{
		"chat_id": "@testForBot1", //
		"text":    "hello world",
	}
	req, err := apiClient.GetRequest("POST", "sendMessage", params)
	common.CheckError(err, 2)
	client := &http.Client{}
	resp, err := client.Do(req)
	common.CheckError(err, 2)
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	common.CheckError(err, 1)
	fmt.Println(string(resBody))
}
