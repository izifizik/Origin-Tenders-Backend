package botService

import "origin-tender-backend/server/internal/domain"

type BotService interface {
	CreateOrder(order domain.Order) error

	GetTgUsers() ([]domain.TelegramUser, error)

	GetSiteUserByName(name string) (domain.User, error)
	CreateSiteUser(user domain.User) error

	GenerateToken(ID string) string
	ProofToken(ID string, token string) error

	CreateTgToken(name string, token string, siteId string) error
}
