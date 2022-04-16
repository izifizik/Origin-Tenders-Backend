package bot

import "go.mongodb.org/mongo-driver/bson/primitive"

type TokenDTO struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type TokenProofDTO struct {
	ID    string
	Token string
}

type OptionsDTO struct {
	TenderID      primitive.ObjectID `json:"tender_id,omitempty"`
	StepPercent   float64            `json:"step_percent,omitempty"`
	MinAutoPrice  float64            `json:"min_auto_price,omitempty"`
	CriticalPrice float64            `json:"critical_price,omitempty"`
}
