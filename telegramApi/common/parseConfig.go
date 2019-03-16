package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// 配置的数据结构
type Config struct {
	BaseUrl  string `json:"baseUrl"`
	ApiToken string `json:"apiToken"`
}

// 解析json中的配置项
func ParseConfig(configPath string) (Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	data := &Config{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalln(err)
	}

	return *data, nil
}
