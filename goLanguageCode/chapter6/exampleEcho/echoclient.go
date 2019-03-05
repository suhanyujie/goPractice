package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	// 获取签名证书
	CA := x509.NewCertPool()
	serverCert, err := ioutil.ReadFile("keys/erp1.com.crt")
	if err != nil {
		log.Fatalln(err)
	}
	CA.AppendCertsFromPEM(serverCert)
	config := tls.Config{RootCAs: CA}
	// 发起连接时，带上证书配置
	conn, err := tls.Dial("tcp", "erp1.com:8001", &config)
	if err != nil {
		log.Fatalf("client:dial:%s \n", err)
	}
	defer conn.Close()
	log.Println("client:connected to:", conn.RemoteAddr())
	state := conn.ConnectionState()
	log.Println("client:handshake:", state.HandshakeComplete)
	log.Println("client:mutual: ", state.NegotiatedProtocolIsMutual)
	message := "Hello \n"
	n, err := io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("client:write:%s", err)
	}
	log.Println("client wrote %q (%d bytes) ", message, n)
	reply := make([]byte, 256)
	n, err = conn.Read(reply)
	log.Printf("client:read %q (%d bytes) ", string(reply[:n]), n)
	log.Println("client exiting...")
}
