package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	Name           string             `bson:"name" json:"name"`
	TendersHistory []Tender           `bson:"tenders_history" json:"tenders_history"`
}
