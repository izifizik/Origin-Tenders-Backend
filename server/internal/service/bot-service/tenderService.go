package botService

import (
	"encoding/json"
	"fmt"
	"origin-tender-backend/server/internal/domain"
	"origin-tender-backend/server/internal/service/wsActions"
)

func (s *service) CreateTender(tender domain.Tender) error {
	return s.repo.CreateTender(tender)
}

func (s *service) UpdateTender(filter interface{}, tender domain.Tender) error {

	data, err := json.Marshal(&tender)
	if err != nil {
		fmt.Println(err)
	}

	wsActions.NotifyAllSession(string(data))
	//tgUsers, err := s.repo.GetTgUsers()

	return s.repo.UpdateTender(filter, tender)
}
