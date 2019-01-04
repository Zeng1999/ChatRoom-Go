package client

type Room struct {
	ID          int
	Name        string
	IsPri       bool
	OnlineUsers []*User
	OwnerId     int
	Messages    []string

	password string
}

func (c *Cli) EnterRoom(uid, roomId int, pass string) error {
	rm := c.GetRoom(roomId)
	if rm == nil {
		return NoSuchRoom
	}

	ur := c.GetUser(uid)
	if ur == nil {
		return NoSuchUser
	}

	if rm.IsPri && pass != rm.password {
		return InvalidPassword
	}

	c.UserEnterRoom(ur, rm)
	return nil
}

func (c *Cli) LeaveRoom(uid, roomId int) error {
	room := c.GetRoom(roomId)
	if room == nil {
		return NoSuchRoom
	}
	user := c.GetUser(uid)
	if user == nil {
		return NoSuchUser
	}
	for i, e := range room.OnlineUsers {
		if e.Uid == user.Uid {
			room.OnlineUsers = append(room.OnlineUsers[:i], room.OnlineUsers[i+1:]...)
			e.CurrentRoomID = -1
			return nil
		}
	}
	return UserNotInRoom
}

func (c *Cli) UserEnterRoom(u *User, r *Room) {
	r.OnlineUsers = append(r.OnlineUsers, u)
	u.CurrentRoomID = r.ID
}

func (c *Cli) CreateRoom(isPri bool, name, pass string, owner int) error {
	if len(name) <= 0 || len(name) > 20 {
		return InvalidName
	}
	if isPri && len(pass) <= 0 {
		return NoPassword
	}
	rm := Room{
		ID:          len(c.rooms) + 1000,
		Name:        name,
		IsPri:       isPri,
		OwnerId:     owner,
		password:    pass,
		OnlineUsers: make([]*User, 0, 10),
		Messages:    make([]string, 0, 10),
	}
	c.rooms = append(c.rooms, &rm)
	return nil
}

func (c *Cli) DeleteRoom(uid, roomId int) error {
	room := c.GetRoom(roomId)
	if room.OwnerId != uid {
		return PermissionDelied
	}
	for i, e := range c.rooms {
		if e.ID == roomId {
			c.rooms = append(c.rooms[:i], c.rooms[i+1:]...)
			return nil
		}
	}
	return NoSuchRoom
}

func (c Cli) GetRoom(roomId int) *Room {
	for _, r := range c.rooms {
		if r.ID == roomId {
			return r
		}
	}
	return nil
}
