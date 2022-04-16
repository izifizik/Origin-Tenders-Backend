package botService

import "origin-tender-backend/server/internal/domain"

func (s *service) CreateTender(tender domain.Tender) error {
	return s.repo.CreateTender(tender)
}
