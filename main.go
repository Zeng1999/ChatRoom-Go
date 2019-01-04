package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var up = websocket.Upgrader{}
var con *websocket.Conn

func main() {
	http.HandleFunc("/", echo)
	go func() {
		for {
			var str string
			fmt.Println("Input Something")
			fmt.Scanln(&str)
			err := con.WriteMessage(websocket.TextMessage, []byte(str))
			if err != nil {
				fmt.Println(err)
				break
			}
			_, mes, err := con.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Get ", string(mes))
		}
	}()
	fmt.Println("Listen 8080 !")
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := up.Upgrade(w, r, nil)
	con = conn
	if err != nil {
		fmt.Fprintln(w, err.Error())
		fmt.Println(err)
		return
	}

	mt, mes, err := con.ReadMessage()
	if err != nil {
		fmt.Fprintln(w, err.Error())
		fmt.Println(err)
		return
	}
	fmt.Println("Get ...")
	fmt.Println(string(mes))

	err = con.WriteMessage(mt, mes)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		fmt.Println(err)
		return
	}
	fmt.Println("Go send OK")
}
