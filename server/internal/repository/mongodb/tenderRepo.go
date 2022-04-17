package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
	"origin-tender-backend/server/internal/service/teleg-bot-service/actions"
)

func (r *repo) CreateTender(tender domain.Tender) error {

	_, err := r.tendersCollection.InsertOne(context.Background(), tender)
	if err != nil {
		fmt.Println("error with create tender: " + err.Error())
		return err
	}
	return nil
}

func (r *repo) UpdateTender(filter interface{}, tender domain.Tender) error {
	_, err := r.tendersCollection.UpdateOne(context.Background(), filter, bson.D{{"$set",
		bson.D{
			{"current_price", tender.CurrentPrice},
			{"status", tender.Status},
		},
	}})

	return err
}

func NotificateTenderChange(tender domain.Tender) {
	var users []domain.TelegramUser

	for _, user := range users {
		actions.NotificateTenderChange(user.UserId, tender)
	}

}

func (r *repo) GetTenderByID(id string) (domain.Tender, error) {
	var tender domain.Tender
	tID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error with ObjectIDFromHex")
		return tender, err
	}
	filter := bson.M{"_id": tID}

	err = r.tendersCollection.FindOne(context.Background(), filter).Decode(&tender)
	if err != nil {
		fmt.Println("Error with get tender")
		return tender, err
	}

	return tender, nil
}
