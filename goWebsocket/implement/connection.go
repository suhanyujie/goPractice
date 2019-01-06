package implement

import (
	"github.com/gorilla/websocket"
	"sync"
	"errors"
)

type Connection struct {
	wsConn  *websocket.Conn
	inChan  chan []byte
	outChan chan []byte
	closeChan chan byte
	isClosed bool
	mutex sync.Mutex
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:wsConn,
		inChan:make(chan []byte, 1000),
		outChan:make(chan []byte, 1000),
		closeChan:make(chan byte,1),
		isClosed:false,
	}
	//启动读协程
	go conn.readLoop()
	go conn.writeLoop()
	return
}

// todo 读取消息
func (_this *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-_this.inChan:
	case <-_this.closeChan:
		err = errors.New("connection is closed...\n")
	}


	return data, nil
}

// todo 写消息
func (_this *Connection) WriteMessage(data []byte) (err error) {
	select {
	case _this.outChan <- data:
	case <-_this.closeChan:
		err = errors.New("connection is closed...\n")
	}
	return
}

// todo 关闭连接
func (_this Connection) Close() {
	//Close是线程安全的
	_this.wsConn.Close()
	_this.mutex.Lock()
	if !_this.isClosed {
		close(_this.closeChan)
		_this.isClosed = true
	}
	_this.mutex.Unlock()
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
		select {

		//inChan是否有空闲位置，如果没有，则阻塞
		case _this.inChan <- data:
		case <-_this.closeChan:
			goto ERR
		}

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
		select {
		case data = <-_this.outChan:
		case <-_this.closeChan:
			goto ERR
		}

		if err = _this.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	_this.Close()
}
