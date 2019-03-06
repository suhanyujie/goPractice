package main

import "practice/telegramApi/apiImpl"

// https://core.telegram.org/bots/api#available-methods
func main() {
	//fmt.Println(configData.BaseUrl)
	apiImpl.SendMessage()
}
