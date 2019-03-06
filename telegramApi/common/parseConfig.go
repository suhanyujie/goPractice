package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	BaseUrl  string `json:"baseUrl"`
	ApiToken string `json:"apiToken"`
}

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
