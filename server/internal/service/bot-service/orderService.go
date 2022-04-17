package botService

import "origin-tender-backend/server/internal/domain"

func (s *service) CreateOrder(order domain.Order) error {
	return s.repo.CreateOrder(order)
}

func (s *service) GetTenderOrders(tenderId string) ([]domain.Order, error) {
	return s.repo.GetTenderOrders(tenderId)
}
