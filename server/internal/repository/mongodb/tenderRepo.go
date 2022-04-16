package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
	"time"
)

func (r *repo) CreateTender(tender domain.Tender) {
	_, err := r.proofTokenCollection.InsertOne(context.Background(), bson.D{
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
	return
	// TODO: do stuff, call bot events
}

func (r *repo) GetTenderByID(id string) domain.Tender {
	tID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error with ObjectIDFromHex")
		return domain.Tender{}
	}
	filter := bson.M{"_id": tID}
	var tender domain.Tender
	err = r.tendersCollection.FindOne(context.Background(), filter).Decode(&tender)
	if err != nil {
		fmt.Println("Error with get tender")
		return domain.Tender{}
	}
	return domain.Tender{}
}
