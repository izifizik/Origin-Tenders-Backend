package bot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"origin-tender-backend/server/internal/domain/wsModels"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *handler) initRead(c *gin.Context) (*websocket.Conn, error) {
	upgrader.CheckOrigin(c.Request)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error with ws connect: " + err.Error())
		return nil, err
	}

	return conn, nil
}

func (h *handler) Notify(c *gin.Context) {

	userId := c.Param("id")

	conn, err := h.initRead(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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

func (h *handler) Bets(c *gin.Context) {

	userId := c.Param("id")

	conn, err := h.initRead(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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

func (h *handler) Session(c *gin.Context) {
	userId := c.Param("id")

	conn, err := h.initRead(c)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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
