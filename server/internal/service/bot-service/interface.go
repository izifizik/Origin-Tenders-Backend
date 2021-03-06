package botService

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
)

type BotService interface {
	CreateOrder(order domain.Order) error
	UpdateTender(tenderId string, tender domain.Tender) error
	GetTenderByID(id string) (domain.Tender, error)

	GetTgUsers() ([]domain.TelegramUser, error)

	GetSiteUserByName(name string) (domain.User, error)

	CreateTgToken(name string, token string, siteId primitive.ObjectID) error

	CreateTender(tender domain.Tender) error
	GetTenderOrders(tenderId string) ([]domain.Order, error)

	BotSetup(id, tenderID, alg, tpe string, procent, minimal, critical float64, isApprove bool)
}
