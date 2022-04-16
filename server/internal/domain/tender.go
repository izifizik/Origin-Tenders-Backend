package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tender struct {
	ID                 primitive.ObjectID
	TimeEnd            time.Time
	Name               string
	Description        string
	Filters            []string
	StartPrice         float32
	CurrentPrice       float32
	Status             string
	MinimalStepPercent float32
	MaxStepPercent     float32
}
