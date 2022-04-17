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
	Owner            string    `json:"owner" bson:"owner"`
	Filters          []string  `json:"filters" bson:"filters"`
	StartPrice       float64   `json:"start_price" bson:"start_price"`
	CurrentPrice     float64   `json:"current_price" bson:"current_price"`
	Status           string    `json:"status" bson:"status"`
	StepPercent      float64   `json:"step_percent" bson:"step_percent"`
}
