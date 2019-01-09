package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"GoApp/HelloGo/GoApp"
	"fmt"
)

var comm *tars.Communicator

// 入口函数
func main() {
	comm = tars.NewCommunicator()
	obj := "TarsGo.HelloGO.SayHelloObj@tcp -h 172.17.0.3 -p 11001"
	app := new(GoApp.SayHello)
	comm.StringToProxy(obj, app)
	var resInt *int32
	ret, err := app.Add(12, 534, resInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ret:%d,result:%d \n", ret, resInt)
}
