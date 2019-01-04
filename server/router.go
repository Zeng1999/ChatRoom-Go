package server

import (
	"fmt"
	"log"
	"net/http"

	"../client"
)

var DefaultClient *client.Cli

func init() {
	cli, err := client.New()
	if err != nil {
		panic(err)
	}
	DefaultClient = cli
}

// Router
func RouteAndListen(port int) {
	http.HandleFunc("/sign", SignHandler)
	http.HandleFunc("/room", RoomHandler)
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/send", SendHandler)
	http.HandleFunc("/info", InfoHandler)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

/*
{
	"operation": {
		"sign": [
			"in",
			"up"
		],
		"room": [
			"create",
			"delete",
			"enter",
			"leave"
		],
		"send": [
			"sendMes"
		],
		"info": [
			"number"
		]
	}
}*/
