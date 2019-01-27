package client

import (
	"time"

	"github.com/Zeng1999/ChatRoom-Go/tools"
	"github.com/gorilla/websocket"
)

type User struct {
	Uid           int
	Name          string
	CurrentRoomID int
	Conn          *websocket.Conn
	MessageQueue  *tools.Queue

	pass string
}

type empty struct{}

func (u *User) routine(stop <-chan empty, errs chan<- error) {
	var err error

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stop:
			break
		case <-ticker.C:
			for !u.MessageQueue.IsEmpty() {
				err = u.Conn.WriteMessage(websocket.TextMessage, []byte(u.MessageQueue.DeQueue()))
				if err != nil {
					errs <- err
				}
			}
		}
	}
}

func (c *Cli) GetUser(uid int) *User {
	for _, u := range c.onlines {
		if u.Uid == uid {
			return u
		}
	}
	return nil
}
