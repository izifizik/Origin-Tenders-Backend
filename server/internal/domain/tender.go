package domain

import "time"

type Tender struct {
	Name        string
	TimeEnd     time.Time
	Description string
	WorkType    string // ObjectId?
	//workType хочу в отдельную сущность
	StartPrice         float32
	CurrentPrice       float32
	Status             string
	MinimalStepPercent float32
	MaxStepPercent     float32
}
