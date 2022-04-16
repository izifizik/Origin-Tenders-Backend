package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tender struct {
	WorkType string // ObjectId?
	//workType хочу в отдельную сущность

	MinimalStepPercent float32
	MaxStepPercent     float32

	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	TimeEnd          time.Time          `json:"timeEnd" bson:"time_end"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description" bson:"description"`
	Filters          []string
	StartPrice       float32 `json:"startPrice" bson:"start_price"`
	CurrentPrice     float32 `json:"currentPrice" bson:"current_price"`
	Status           string  `json:"status" bson:"status"`
	StepPercent      float32
	ShortDescription string `json:"shortDescription"`
}
