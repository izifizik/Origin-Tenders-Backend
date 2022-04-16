package botService

import (
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
)

type BotService interface {
	CreateOrder(order domain.Order) error

	GetSiteUserByName(name string) (domain.User, error)
	CreateSiteUser(user domain.User) error

	GenerateToken(ID string) string
	ProofToken(ID string, token string) error

	CreateTgToken(name string, token string, siteId primitive.ObjectID) error

	SentNotification(conn *websocket.Conn)
}
