package main

import (
	"net"
	"log"
	"fmt"
	"time"
	"io/ioutil"
	"bufio"
	"strings"
)

func main() {
	server()

	//c1 := make(chan string)
	////c2 := make(chan string)
	//go testDail(c1);
	//fmt.Println(<-c1)
}

func testDail(c1 chan string) {
	//创建tcp连接，此处的超时是发起请求后开始计算时间
	conn, err := net.DialTimeout("tcp", "www.baidu.com:80", time.Second*2)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	//向连接中写入请求数据
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	//读取响应
	//status, err := bufio.NewReader(conn).ReadString('\n')
	//if err != nil {
	//	log.Println(err)
	//}
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))

	c1 <- string(data)
}

// todo
func server() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn);
	}
}

func handleConnection(conn net.Conn) {
	for {
		status, err := bufio.NewReader(conn).ReadString('\n');
		if err != nil {
			fmt.Println(err)
			return
		}
		status = strings.TrimSpace(status)
		fmt.Println(len(strings.TrimSpace(status)))
		fmt.Println(status)
		switch status {
		case "123":
			conn.Write([]byte("response is 123...\n"))
		case "456":
			conn.Write([]byte("response is 456...\n"))
		default:
			conn.Write([]byte("response is default value...\n"))

		}
	}
}

/**




## 参考地址 https://studygolang.com/pkgdoc
 */
