package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
Sign Json Data
{
	"oper": "in",
	"user": {
		"name": "",
		"pass": ""
	}
}
*/
func SignHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(200)
		rMes(w, true, "Please Post Data !")
		return
	}
	var signData SignData
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rMes(w, false, err.Error())
		return
	}
	err = json.Unmarshal(data, &signData)
	if err != nil {
		rMes(w, false, err.Error())
		return
	}
	switch signData.Oper {
	case "in":
		err = DefaultClient.SignIn(w, r, signData.User.Name, signData.User.Pass)
	case "up":
		err = DefaultClient.SignUp(signData.User.Name, signData.User.Pass)
	default:
		err = errors.New("Unknown Operation !")
	}
	if err != nil {
		rMes(w, false, err.Error())
		return
	}
	rMes(w, true, fmt.Sprintf("Sign %s Success !", signData.Oper))
}

/*
Room Json Data
{
	"oper": "",
	"user": {
		"name": "",
		"pass": ""
	},
	"room": {
		"rid": 123,
		"name": "",
		"ispri": false,
		"pass": ""
	}
}
*/
func RoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(200)
		rMes(w, true, "Please Post Data !")
		return
	}
	var roomData RoomData
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rMes(w, false, err.Error())
		return
	}
	err = json.Unmarshal(data, &roomData)
	if err != nil {
		rMes(w, false, err.Error())
		return
	}

	switch roomData.Oper {
	case "create":
		err = DefaultClient.CreateRoom(roomData.Room.IsPri, roomData.Room.Name, roomData.Room.Pass, roomData.User.Uid)
	case "delete":
		err = DefaultClient.DeleteRoom(roomData.User.Uid, roomData.Room.Rid)
	case "enter":
		err = DefaultClient.EnterRoom(roomData.User.Uid, roomData.Room.Rid, roomData.Room.Pass)
	case "leave":
		err = DefaultClient.LeaveRoom(roomData.User.Uid, roomData.Room.Rid)
	default:
		err = errors.New("Unknown Operation !")
	}
	if err != nil {
		rMes(w, false, err.Error())
		return
	}
	rMes(w, true, fmt.Sprintf("%s Success !", roomData.Oper))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	rMes(w, true, "Just An Open Interface")
}

/*
Send Json Data
{
	"oper": "sendMes",
	"uid": 123,
	"mess": "asdasd"
}
*/
func SendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(200)
		rMes(w, true, "Please Post Data !")
		return
	}
	var sendData SendData
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rMes(w, false, err.Error())
		return
	}
	err = json.Unmarshal(data, &sendData)
	if err != nil {
		rMes(w, false, err.Error())
		return
	}

	switch sendData.Oper {
	case "sendMes":
		err = DefaultClient.SendMessage(sendData.Uid, websocket.TextMessage, []byte(sendData.Mess))
	default:
		err = errors.New("Unknown Operation !")
	}
	if err != nil {
		rMes(w, false, err.Error())
		return
	}
	rMes(w, true, "Send Success !")
}

/*
Info Json Data
{
	"oper": "number",
	"type": "room",
	"id": 123
}
*/
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(200)
		rMes(w, true, "Please Post Data !")
		return
	}
	var infoData InfoData
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rMes(w, false, err.Error())
		return
	}
	err = json.Unmarshal(data, &infoData)
	if err != nil {
		rMes(w, false, err.Error())
		return
	}

	switch infoData.Oper {
	case "number":
		if infoData.RoomOrUser == "room" {
		}
	default:
	}
}

func rMes(w http.ResponseWriter, ok bool, tip string) {
	data, err := json.Marshal(Mes{
		Ok:  ok,
		Tip: tip,
	})
	if err != nil {
		fmt.Println("Json Marshal Error !", err)
		fmt.Fprintln(w, "Data Error !")
		return
	}
	fmt.Fprintf(w, string(data))
}
