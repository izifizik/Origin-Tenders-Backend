package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID
	Name string
	Notifications
	Filters        []string
	TendersHistory []Tender
}

type Notifications struct {
	TgID string
}
