package rule

func GetAreaRule() string {
	areaRule := "#list"
	//areaRule = base64.StdEncoding.EncodeToString([]byte(areaRule))
	return areaRule
}

func GetItemRule() map[string]string {
	var ruleArr = map[string]string{
		"0": "dd a",
	}
	for index, con := range ruleArr {
		//newRule := base64.StdEncoding.EncodeToString([]byte(con))
		ruleArr[index] = string(con)
	}
	return ruleArr
}
