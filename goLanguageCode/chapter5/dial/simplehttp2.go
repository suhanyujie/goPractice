package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"practice/goLanguageCode/chapter5/dial/common"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	//解析host和端口号
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	common.CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	common.CheckError(err)
	_, err = conn.Write([]byte("HEAD / HTTP1.0\r\n\r\n"))
	common.CheckError(err)
	result, err := ioutil.ReadAll(conn)
	common.CheckError(err)
	fmt.Println(string(result))

	os.Exit(0)
}
