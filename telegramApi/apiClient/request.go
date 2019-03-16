package apiClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"practice/telegramApi/common"
)

var configData = common.Config{}

func init() {
	tmpConfigData, err := common.ParseConfig("config.json")
	if err != nil {
		common.CheckError(err, 2)
	}
	configData = tmpConfigData
}

func GetRequest(method, apiName string, params map[string]string) (*http.Request, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(body)
	url := getApiUrl(apiName)
	fmt.Println(url)
	request, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-type", "application/json")

	return request, nil
}

func getApiUrl(apiName string) string {
	return configData.BaseUrl + "/bot" + configData.ApiToken + "/" + apiName
}
