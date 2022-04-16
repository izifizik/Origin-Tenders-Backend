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

	ID           primitive.ObjectID
	TimeEnd      time.Time
	Name         string
	Description  string
	Filters      []string
	StartPrice   float32
	CurrentPrice float32
	Status       string
	StepPercent  float32
}
