package main

import (
	"crypto/rand"
	"crypto/tls"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("keys/erp1.com.crt", "keys/erp1.com.key")
	if err != nil {
		log.Fatal(err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader
	var port = "8001"
	service := "127.0.0.1:" + port
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server listening in port:", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("server:accept from ", conn.RemoteAddr())
		go HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	for {
		log.Println("server:conn is handling")
		_, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		n, err := conn.Write(buf)
		log.Println("the server has write ", n, " byte data to client!")
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("the conn of the client will cloed...")
}
