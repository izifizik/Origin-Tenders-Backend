package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
)

func (r *repo) CreateTender(tender domain.Tender) error {

	_, err := r.tendersCollection.InsertOne(context.Background(), tender)
	if err != nil {
		fmt.Println("error with create tender: " + err.Error())
		return err
	}
	return nil
}

func (r *repo) UpdateTender(tenderId string, tender domain.Tender) error {
	id, _ := primitive.ObjectIDFromHex(tenderId)
	_, err := r.tendersCollection.UpdateOne(context.Background(),
		bson.M{"_id": id},
		bson.D{{"$set",
			bson.D{
				{"current_price", tender.CurrentPrice},
				{"status", tender.Status},
			},
		}})

	return err
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
