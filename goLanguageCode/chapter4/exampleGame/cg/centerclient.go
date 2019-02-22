package cg

import (
	"encoding/json"
	"errors"
	"practice/goLanguageCode/chapter4/exampleGame/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}
	resp, err := client.Call("addPlayer", string(b))
	if err == nil || resp.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("removePlayer", name)
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, _ := client.Call("listPlayer", params)
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	if err != nil {
		ps = nil
	}
	return
}

func (client *CenterClient) BroadCast(message string) error {
	m := &Message{Content: message}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp, _ := client.Call("broadCast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}
