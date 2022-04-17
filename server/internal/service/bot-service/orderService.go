package botService

import (
	"fmt"
	"origin-tender-backend/server/internal/domain"
)

func (s *service) CreateOrder(order domain.Order) error {

	tender, err := s.repo.GetTenderByID(order.TenderId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tender.CurrentPrice = order.Price

	err = s.UpdateTender(order.TenderId, tender)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return s.repo.CreateOrder(order)
}

func (s *service) GetTenderOrders(tenderId string) ([]domain.Order, error) {
	return s.repo.GetTenderOrders(tenderId)
}
