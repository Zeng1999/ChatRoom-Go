package client

import (
	"net/http"

	"github.com/Zeng1999/ChatRoom-Go/tools"
)

func (c *Cli) SignUp(name, pass string) error {
	alreadyHave, err := c.db.CheckExists(name)
	if err != nil {
		return err
	}
	if alreadyHave {
		return NameExists
	}
	pass_md5, err := tools.MD5(pass)
	if err != nil {
		return err
	}
	return c.db.AddUser(name, pass_md5)
}

func (c *Cli) SignIn(w http.ResponseWriter, r *http.Request, name, pass string) error {
	ur, err := c.db.GetUser(name)
	if err != nil {
		return err
	}
	pass_md5, err := tools.MD5(pass)
	if err != nil {
		return err
	}
	if ur.Pass != pass_md5 {
		return InvalidPassword
	}
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	u := User{
		Uid:           ur.Uid,
		Name:          name,
		CurrentRoomID: -1,
		pass:          pass,
		Conn:          conn,
		MessageQueue:  tools.NewQueue(),
	}
	c.onlines = append(c.onlines, &u)
	return nil
}
