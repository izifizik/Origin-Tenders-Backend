package websocket

import "github.com/gorilla/websocket"

type UserClient struct {
	Name string
	conn *websocket.Conn
}

type Action struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
