package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string
	Bot            []Bot
	Filters        []Filter
	TendersHistory []Tender
}

type Bot struct {
	TenderID      primitive.ObjectID
	StepPercent   float64 `bson:"step_percent"`
	CriticalPrice float64
	IsNeedApprove bool
}

type Filter struct {
}
