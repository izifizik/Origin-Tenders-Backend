package botService

import (
	"origin-tender-backend/server/internal/domain"
)

func (s *service) GetTgUsers() ([]domain.TelegramUser, error) {
	return s.repo.GetTgUsers()
}
