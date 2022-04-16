package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
	"time"
)

func (r *repo) CreateTender(tender domain.Tender) error {
	_, err := r.tendersCollection.InsertOne(context.Background(), bson.D{
		{"timeStamp", time.Now()},
		{"Name", tender.Name},
		{"Description", tender.Description},
		{"StartPrice", tender.StartPrice},
		{"StartPrice", tender.StartPrice},
		{"status", "open"},
	})
	if err != nil {
		fmt.Println("error with create tender: " + err.Error())
	}
	return err
	// TODO: do stuff, call bot events
}

func (r *repo) GetTenderByID(id string) domain.Tender {
	var tender domain.Tender
	tID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error with ObjectIDFromHex")
		return tender
	}
	filter := bson.M{"_id": tID}

	err = r.tendersCollection.FindOne(context.Background(), filter).Decode(&tender)
	if err != nil {
		fmt.Println("Error with get tender")
		return domain.Tender{}
	}
	return domain.Tender{}
}
