package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"origin-tender-backend/server/internal/domain"
	"time"
)

func (r repo) CreateTender(tender domain.Tender) error {
	_, err := r.proofTokenCollection.InsertOne(context.Background(), bson.D{
		{"timeStamp", time.Now()},
		{"MinimalStepPercent", tender.MinimalStepPercent},
		{"MaxStepPercent", tender.MaxStepPercent},
		{"Name", tender.Name},
		{"Description", tender.Description},
		{"StartPrice", tender.StartPrice},
		{"StartPrice", tender.StartPrice},
		{"status", "open"},
	})

	// TODO: do stuff, call bot events

	return err
}

func (r repo) GetTenders() []domain.Tender {

	return nil
}






