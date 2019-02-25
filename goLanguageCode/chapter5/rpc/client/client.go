package main

import (
	"fmt"
	"log"
	"net/rpc"
	"practice/goLanguageCode/chapter5/rpc/server"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Fatal("dialing error:", err)

	}
	args := &server.Args{
		12, 10,
	}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	////异步调用
	//quotient := new(server.Quotient)
	//divCall := client.Go("Arith.Divide", args, &quotient, nil)
	//replyCall := <-divCall.Done
	//fmt.Println(replyCall)

	fmt.Println("the result is:", reply)
}
