package main

import (
	"fmt"
	"net"
	"os"
	"practice/goLanguageCode/chapter5/dial/common"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	common.CheckError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0 \r\n\r\n"))
	common.CheckError(err)
	result, err := common.ReadFully(conn)
	common.CheckError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
