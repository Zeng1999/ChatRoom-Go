package client

import (
	"errors"

	mq "github.com/Zeng1999/ChatRoom-Go/mysql"
	"github.com/gorilla/websocket"
)

type Cli struct {
	rooms    []*Room
	onlines  []*User
	db       mq.DB
	upgrader websocket.Upgrader
}

var (
	NoSuchRoom       = errors.New("No Such Room !")
	NoSuchUser       = errors.New("No Such User !")
	InvalidName      = errors.New("Invalid Name !")
	InvalidPassword  = errors.New("Invalid Password !")
	UserNotInRoom    = errors.New("User Not In This Room !")
	NameExists       = errors.New("Name Exists !")
	NoPassword       = errors.New("Need Password !")
	PermissionDelied = errors.New("Permission Delied !")
)

func New() (*Cli, error) {
	dc, err := mq.ReadConfig()
	if err != nil {
		return nil, err
	}

	db, err := mq.New(*dc)
	if err != nil {
		panic(err)
	}

	cli := Cli{
		rooms:    make([]*Room, 0, 5),
		onlines:  make([]*User, 0, 10),
		db:       db,
		upgrader: websocket.Upgrader{},
	}
	return &cli, nil
}

func (c *Cli) SendToRoom(roomId int, messageType int, data []byte) error {
	rm := c.GetRoom(roomId)
	if rm == nil {
		return NoSuchRoom
	}

	var err error
	for _, e := range rm.OnlineUsers {
		err = e.Conn.WriteMessage(messageType, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cli) BroadcastMessage(messageType int, data []byte) error {
	var err error
	for _, e := range c.onlines {
		err = e.Conn.WriteMessage(messageType, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cli) SendMessage(uid int, messageType int, data []byte) error {
	ur := c.GetUser(uid)
	if ur == nil {
		return NoSuchUser
	}
	roomId := ur.CurrentRoomID
	return c.SendToRoom(roomId, messageType, data)
}
