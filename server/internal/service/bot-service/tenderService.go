package botService

import "origin-tender-backend/server/internal/domain"

func (s *service) CreateTender(tender domain.Tender) {
	s.repo.CreateTender(tender)
}

func (s *service) GetTenderByID(id string) domain.Tender {
	//s.repo.GetTenderByID(id)
	return domain.Tender{}
}
