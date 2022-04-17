package wsModels

import "github.com/gorilla/websocket"

var UserConnections = make(map[string]WsConnections)

type WsConnections struct {
	Bets         *websocket.Conn
	Session      *websocket.Conn
	Notification *websocket.Conn
}
