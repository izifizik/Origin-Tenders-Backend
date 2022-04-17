package botService

import "origin-tender-backend/server/internal/domain"

func (s *service) CreateOrder(order domain.Order) error {
	return s.repo.CreateOrder(order)
}
