package implement

import "github.com/gorilla/websocket"

type Connection struct {
	wsConn  *websocket.Conn
	inChan  chan []byte
	outChan chan []byte
}

func InitConnection(wsConn *websocket.Conn) (*Connection, error) {
	var conn = &Connection{
		wsConn,
		make(chan []byte, 1000),
		make(chan []byte, 1000),
	}
	return
}

// todo 读取消息
func (_this *Connection) ReadMessage() (data []byte, err error) {
	data = <-_this.inChan
	return data, nil
}

// todo 写消息
func (_this *Connection) WriteMessage(data []byte) (error) {
	_this.outChan <- data
	return nil
}

// todo 关闭连接
func (_this Connection) Close() {
	//Close是线程安全的
	_this.wsConn.Close()
}

// todo 读取连接中的消息
func (_this *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = _this.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		_this.inChan <- data
	}
ERR:
	_this.Close()
}

// todo 向连接中写入数据
func (_this *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		data = <-_this.outChan
		if err = _this.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	_this.Close()
}
