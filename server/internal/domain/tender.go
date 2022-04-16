package domain

import (
	"time"
)

type Tender struct {
	ID               string    `json:"_id" bson:"_id"`
	TimeEnd          time.Time `json:"timeEnd" bson:"time_end"`
	Name             string    `json:"name" bson:"name"`
	Description      string    `json:"description" bson:"description"`
	ShortDescription string    `json:"short_description" bson:"short_description"`
	Filters          []string
	StartPrice       float32 `json:"startPrice" bson:"start_price"`
	CurrentPrice     float32 `json:"currentPrice" bson:"current_price"`
	Status           string  `json:"status" bson:"status"`
	StepPercent      float64 `json:"stepPercent"`
}
