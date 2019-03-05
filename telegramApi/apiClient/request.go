package apiClient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"practice/telegramApi/config"
)

func GetRequest(method, apiName string, params map[string]string) (*http.Request, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(body)
	url := getApiUrl(apiName)
	request, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-type", "application/json")

	return request, nil
}

func getApiUrl(apiName string) string {
	return config.ApiUrl + "/bot" + config.ApiToken + "/" + apiName
}
