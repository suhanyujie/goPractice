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
	obj := "GoApp.HelloGo.SayHelloObj"
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 172.17.0.3 -p 17890")
	app := new(GoApp.SayHello)
	comm.StringToProxy(obj, app)
	var i1,i2,resInt int32
	i1 = 12
	i2 = 534
	ret, err := app.Add(i1, i2, &resInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ret:%d,result:%d \n", ret, resInt)
}
