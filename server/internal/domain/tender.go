package domain


import "time"
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tender struct {
	Name        string
	TimeEnd     time.Time
	Description string
	WorkType    string // ObjectId?
	//workType хочу в отдельную сущность

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
