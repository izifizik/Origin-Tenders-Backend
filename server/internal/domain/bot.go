package domain



type Bot struct {
}

//type Bot struct {
//
//}


//type AccountID

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bot struct {
	UserID  primitive.ObjectID
	Filters []string
	Options []Options
}

type Options struct {
	TenderID      primitive.ObjectID
	StepPercent   float64
	MinAutoPrice  float64
	CriticalPrice float64
}

