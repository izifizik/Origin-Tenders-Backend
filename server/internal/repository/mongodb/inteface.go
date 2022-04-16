package mongodb

import (
	"origin-tender-backend/server/internal/domain"
)

type Repository interface {

	GetTgUser(name string) (domain.TelegramUser, string, error)
	CreateNewTgUser(id int64, name string, token string) (domain.TelegramUser, string, error)

	SaveToken(ID string, token string) error
	ProofToken(ID string, token string) error

}
