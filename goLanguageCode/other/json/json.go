package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var str1 = "{\"0\":\"who are u\"}"
	data := []string{}
	err := json.Unmarshal([]byte(str1), &data)
	fmt.Println(err)
	fmt.Println(data)

}
