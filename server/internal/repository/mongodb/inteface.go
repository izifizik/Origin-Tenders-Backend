package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
)

type Repository interface {
	CreateTender(tender domain.Tender) error        // ok
	GetTenderByID(id string) (domain.Tender, error) // ok
	GetTenderOrders(tenderId string) ([]domain.Order, error)
	UpdateTender(filter interface{}, tender domain.Tender) error

	CreateOrder(order domain.Order) error //ok

	GetTgUsers() ([]domain.TelegramUser, error) //ok

	GetSiteUserByName(name string) (domain.User, error) //ok
	CreateSiteUser(user domain.User) error              //ok

	GetUserByTgId(id int64) (domain.TelegramUser, error) //ok
	GetTgUser(name string) (domain.TelegramUser, string, error)
	CreateNewTgUser(id int64, name string, token string) (domain.TelegramUser, string, error)
	UpdateUserStateById(id string, state string) error
	SetUserNameById(idStr string, name string) error
	SetSiteId(siteId string, idStr string) error

	CreateToken(name string, token string, siteId primitive.ObjectID) error
	ApproveProofToken(name string, token string) (string, error)

	SaveToken(ID string, token string) error
	ProofToken(ID string, token string) error
}
