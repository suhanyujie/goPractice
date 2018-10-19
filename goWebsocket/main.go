package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"io/ioutil"
	"os"
	"fmt"
)

/**
## 过程
1.实现一个http接口



*/

// 入口函数
func main() {
	http.HandleFunc("/index", HandleFunc1)
	http.HandleFunc("/staticfile/", HandleStatic)
	http.HandleFunc("/ws", WsHandler)
	http.ListenAndServe("0.0.0.0:3031", nil);
}

var upgrader = websocket.Upgrader{
	//允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var WsHandler = func(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err  error
		//msgType int
		data []byte
	)
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// websocket conn
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			log.Println(err)
			break
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println(err)
			break
		}
	}
	conn.Close()
}

var HandleFunc1 = func(w http.ResponseWriter, r *http.Request) {
	htmlstring, err := ioutil.ReadFile("/www/2017/go/src/practice/goWebsocket/static/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("receive conn...")
	w.Write(htmlstring)
}

var HandleStatic = func (w http.ResponseWriter,r *http.Request) {
	var (
		wd string
		err error
	)
	if wd,err=os.Getwd();err!=nil {
		log.Println(err)
		return
	}
	var filePath = fmt.Sprintf("%s/goWebsocket/static/",wd)
	http.StripPrefix("/staticfile/static/", http.FileServer(http.Dir(filePath))).ServeHTTP(w,r)
}
