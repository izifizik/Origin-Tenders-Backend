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
	tenderID, err := primitive.ObjectIDFromHex(tender.ID)
	if err != nil {
		return err
	}
	_, err = r.tendersCollection.InsertOne(context.Background(), bson.D{
		{"_id", tenderID},
		{"time_end", time.Now().Add(time.Hour * 24)},
		{"name", tender.Name},
		{"description", tender.Description},
		{"short_description", tender.ShortDescription},
		{"owner", tender.Owner},
		{"filters", tender.Filters},
		{"StartPrice", tender.StartPrice},
		{"current_price", tender.StartPrice},
		{"status", "Активно"},
		{"step_percent", tender.StepPercent},
	})
	if err != nil {
		fmt.Println("error with create tender: " + err.Error())
		return err
	}
	return nil
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
