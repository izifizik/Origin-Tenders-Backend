package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string
	TendersHistory []Tender
}

type Bot struct {
	UserID    string
	BotConfig BotConfig
}

type BotConfig struct {
	Alg       string
	TenderID  string
	Type      string
	Procent   float64
	Minimal   float64
	Critical  float64
	IsApprove bool
}
