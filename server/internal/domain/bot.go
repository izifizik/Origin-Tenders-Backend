package domain

type Bot struct {
	UserID    string    `json:"userID" bson:"userID"`
	BotConfig BotConfig `json:"botConfig" bson:"botConfig"`
}

type BotConfig struct {
	Alg       string  `bson:"alg" json:"alg"`
	TenderID  string  `bson:"tenderID" json:"tenderID"`
	Type      string  `bson:"type" json:"type"`
	Procent   float64 `bson:"procent" json:"procent"`
	Minimal   float64 `bson:"minimal" json:"minimal"`
	Critical  float64 `bson:"critical" json:"critical"`
	IsApprove bool    `bson:"isApprove" json:"isApprove"`
}
