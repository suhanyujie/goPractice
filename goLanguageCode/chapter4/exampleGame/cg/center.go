package cg

import (
	"encoding/json"
	"errors"
	"practice/goLanguageCode/chapter4/exampleGame/ipc"
	"strconv"
	"sync"
	"time"
)

type Message struct {
	From    string
	To      string
	Content string
	Time    time.Time
}

type Room struct {
	Id    int
	Desc  string
	Other string
}

type CenterServer struct {
	server  map[string]ipc.IpcServer
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.IpcServer)
	players := make([]*Player, 0)
	return &CenterServer{servers, players, nil, sync.RWMutex{}}
}

func (server *CenterServer) addPlayer(params string) error {
	date1 := strconv.FormatInt(time.Now().Unix(), 10)
	player := NewPlayer(date1)
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.players = append(server.players, player)
	return nil
}

func (server *CenterServer) RemovePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for i, v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 {
				server.players = make([]*Player, 0)
			} else if i == len(server.players)-1 {
				server.players = server.players[:i-1]
			} else if i == 0 {
				server.players = server.players[1:]
			} else {
				server.players = append(server.players[:i-1], server.players[i+1:]...)
			}
		}
	}
	return nil
}

func (server *CenterServer) listPlayer(params string) (players string, err error) {
	server.mutex.RLock()
	defer server.mutex.RUnlock()

	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("No player online.")
	}
	return
}

func (server *CenterServer) broadCast(params string) error {
	var message Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.players) > 0 {
		for _, player := range server.players {
			player.mq <- &message
		}
	} else {
		err = errors.New("No player online.")
	}
	return err
}

func (server *CenterServer) Handle(method, params string) *ipc.Response {
	switch method {
	case "addPlayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removePlayer":
		err := server.RemovePlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listPlayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200", Body: players}
	case "broadCast":
		err := server.broadCast(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{"200", "broadCast successful! \n"}
	default:
		return &ipc.Response{Code: "404", Body: method + ":" + params}
	}
	return nil
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}
