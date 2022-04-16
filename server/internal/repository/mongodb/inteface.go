package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
)

type Repository interface {
	CreateTender(tender domain.Tender) error

	CreateOrder(order domain.Order) error

	GetTgUsers() ([]domain.TelegramUser, error)

	GetSiteUser(objectId string) (domain.User, error)
	GetSiteUserByName(name string) (domain.User, error)
	CreateSiteUser(user domain.User) error

	GetUserByTgId(id int64) (domain.TelegramUser, error)
	GetTgUser(name string) (domain.TelegramUser, string, error)
	CreateNewTgUser(id int64, name string, token string) (domain.TelegramUser, string, error)
	UpdateUserStateById(id string, state string) error
	SetUserNameById(idStr string, name string) error
	SetSiteId(siteId string, idStr string) error

	CreateToken(name string, token string, siteId primitive.ObjectID) error
	ApproveProofToken(name string, token string) (string, error)

	SaveToken(ID string, token string) error
	ProofToken(ID string, token string) error

	CreateBotByID(id, tenderId string, stepPercent, criticalPrice float64, isNeedApprove bool)
}
