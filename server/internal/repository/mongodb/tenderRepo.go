package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"origin-tender-backend/server/internal/domain"
	"time"
)

func (r repo) CreateTender(tender domain.Tender) {
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
	if err != nil {
		fmt.Println("error with create tender: " + err.Error())
	}
	return
	// TODO: do stuff, call bot events
}

func (r repo) GetTender() domain.Tender {

	return domain.Tender{}
}
