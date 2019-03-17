package main

import "practice/telegramApi/apiImpl"

// 尝试发送telegram的消息给机器人
// https://core.telegram.org/bots/api#available-methods
func main() {
	//fmt.Println(configData.BaseUrl)
	// 尝试发送markdown格式的数据
	mdCon := "```\n " +
		"hello world" +
		"\n```"

	apiImpl.SendMessage(mdCon)
}
