package implement

import "github.com/gorilla/websocket"

type Connection struct {
	wsConn *websocket.Conn
	inChan chan []byte
	outChan chan []byte
}

func InitConnection(wsConn *websocket.Conn)(*Connection, error){
	var conn  = &Connection{
		wsConn,
		make(chan []byte,1000),
		make(chan []byte, 1000),
	}
	return
}
