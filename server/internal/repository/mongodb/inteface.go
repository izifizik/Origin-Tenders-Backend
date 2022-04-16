package mongodb

import (
	"origin-tender-backend/server/internal/domain"
)

type Repository interface {
	GetTgUser(name string) (domain.TelegramUser, string, error)
	CreateNewTgUser(id int64, name string, token string) (domain.TelegramUser, string, error)
}
