package bot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"origin-tender-backend/websocket/wsModels"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//func (h *handler) NotificationWS(c *gin.Context) {
//	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		log.Println(err.Error())
//		return
//	}
//	// go + уведы в тг бота
//	h.botService.SentNotification(conn)
//}

func initRead(c *gin.Context) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}

	return conn, nil
}

func Notification(c *gin.Context) {

	userId := c.Param("id")

	conn, err := initRead(c)
	if err != nil {
		fmt.Println(err)
	}

	if userId != "" {
		wsModels.UserConnections[userId] = wsModels.WsConnections{
			Bets:         wsModels.UserConnections[userId].Bets,
			Session:      wsModels.UserConnections[userId].Session,
			Notification: conn,
		}
		conn.WriteMessage(1, []byte("ok"))
	}

}

func Notify(c *gin.Context) {

	userId := c.Param("id")

	conn, err := initRead(c)
	if err != nil {
		fmt.Println(err)
	}

	if userId != "" {
		wsModels.UserConnections[userId] = wsModels.WsConnections{
			Bets:         wsModels.UserConnections[userId].Bets,
			Session:      wsModels.UserConnections[userId].Session,
			Notification: conn,
		}

		conn.WriteMessage(1, []byte("ok"))
	}

}

func Bets(c *gin.Context) {

	userId := c.Param("id")

	conn, err := initRead(c)
	if err != nil {
		fmt.Println(err)
	}

	if userId != "" {
		wsModels.UserConnections[userId] = wsModels.WsConnections{
			Bets:         conn,
			Session:      wsModels.UserConnections[userId].Session,
			Notification: wsModels.UserConnections[userId].Notification,
		}

		conn.WriteMessage(1, []byte("ok"))
	}

}

func Session(c *gin.Context) {

	userId := c.Param("id")

	conn, err := initRead(c)
	if err != nil {
		fmt.Println(err)
	}

	if userId != "" {
		wsModels.UserConnections[userId] = wsModels.WsConnections{
			Bets:         wsModels.UserConnections[userId].Bets,
			Session:      conn,
			Notification: wsModels.UserConnections[userId].Notification,
		}

		conn.WriteMessage(1, []byte("ok"))
	}

}
