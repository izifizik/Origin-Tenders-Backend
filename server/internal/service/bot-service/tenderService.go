package botService

import (
	"encoding/json"
	"fmt"
	"origin-tender-backend/server/internal/domain"
	"origin-tender-backend/server/internal/service/teleg-bot-service/actions"
	"origin-tender-backend/server/internal/service/wsActions"
)

func (s *service) CreateTender(tender domain.Tender) error {
	return s.repo.CreateTender(tender)
}

func (s *service) GetTenderByID(id string) (domain.Tender, error) {
	return s.repo.GetTenderByID(id)
}

func (s *service) UpdateTender(tenderId string, tender domain.Tender) error {

	data, err := json.Marshal(&tender)
	if err != nil {
		fmt.Println(err)
	}

	wsActions.NotifyAllSession(string(data))
	tgUsers, err := s.repo.GetTgUsers()
	if err != nil {
		fmt.Println(err)
	}
	NotificateTenderChange(tgUsers, tender)

	return s.repo.UpdateTender(tenderId, tender)
}

func NotificateTenderChange(users []domain.TelegramUser, tender domain.Tender) {

	for _, user := range users {
		actions.NotificateTenderChange(user.UserId, tender)
	}

}
