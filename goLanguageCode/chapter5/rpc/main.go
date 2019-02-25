package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"practice/goLanguageCode/chapter5/rpc/server"
)

func main() {

	type A1 struct {
		A, B int
	}

	var m1 = map[A1]string{
		(A1{1, 2}): "who",
	}
	fmt.Println(m1)
	os.Exit(0)

	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	e := http.ListenAndServe(":8001", nil)
	if e != nil {
		log.Fatal("listen error:", e)
	}
}
