package wsModels

import "github.com/gorilla/websocket"


type


type Connections struct {
	Bets         *websocket.Conn
	Session      *websocket.Conn
	Notification *websocket.Conn
}
