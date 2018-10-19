package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"os"
	"strings"
)

// 入口函数
func main() {
	CrateConn2()
}

/**
1.将标准输入作为输入，将数据传入到程序当中





 */
func CrateConn2() {
	var addr = "127.0.0.1:3130"
	//解析地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		log.Println(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println(err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("hello my client:")
	defer conn.Close()
	for {
		fmt.Print(addr + "> ")
		text,_ := reader.ReadString('\n')
		text = strings.Replace(text,"\n","",-1)
		fmt.Printf("u have input:%s ^^ \n",text)
	}
}

func CreateConn1() {
	conn, err := net.Dial("tcp", "127.0.0.1:3130")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
