package common

import (
	"fmt"
	"net"
)

func CheckIsIp(ip string) bool {
	res := net.ParseIP(ip)
	if res != nil || len(res) > 0 {
		return true
	}
	return false
}

func LookupHost() {
	host := "www.baidu.com"
	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addrs)
}
