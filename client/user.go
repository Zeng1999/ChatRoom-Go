package client

import (
	"github.com/gorilla/websocket"
)

type User struct {
	Uid           int
	Name          string
	CurrentRoomID int
	Conn          *websocket.Conn

	pass string
}

func (c Cli) GetUser(uid int) *User {
	for _, u := range c.onlines {
		if u.Uid == uid {
			return u
		}
	}
	return nil
}
