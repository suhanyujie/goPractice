package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage:", os.Args[0], "host")
	//	os.Exit(1)
	//}
	//service := os.Args[1]
	service := "www.baidu.com"
	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)
	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)
	_, err = conn.Write(msg[0:len])
	checkError(err)
	n, err := conn.Read(msg[0:])
	fmt.Println(n)
	checkError(err)
	fmt.Println("Got response...")
	fmt.Println("Got response...")
	fmt.Println("Got response...")
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Squence matches")
	}
	os.Exit(0)
}

// 书中示例代码实现
func checkSum1(msg []byte) uint16 {
	sum := 0
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

// 正常运行的实现  https://moekr.com/article/icmp-ping-in-go.html
func checkSum(data []byte) uint16 {
	var (
		sum    uint32
		length = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += sum >> 16
	return uint16(^sum)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1001)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}
