package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                primitive.ObjectID
	Name              string
	TenderParticipant []primitive.ObjectID
	Filters           []string
	TendersHistory    []Tender
}

type TenderParticipant struct {
}
