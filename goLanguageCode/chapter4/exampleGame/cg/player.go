package cg

import (
	"fmt"
)

type Player struct {
	Name  string "name"
	Level int    "level"
	Exp   int
	Room  int
	mq    chan *Message
}

func NewPlayer(name string) *Player {
	m := make(chan *Message, 100)
	p := &Player{
		name,
		1,
		0,
		1,
		m,
	}
	go func(player *Player) {
		for {
			msg := <-player.mq
			fmt.Println(p.Name, "received message:", msg.Content)
		}
	}(p)

	return p
}
