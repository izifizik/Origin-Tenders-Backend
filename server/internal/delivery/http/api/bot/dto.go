package bot

type TokenDTO struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type TokenProofDTO struct {
	ID    string
	Token string
}

type BotSetupDTO struct {
	UserID    string  `json:"user_id,omitempty"`
	TenderID  string  `json:"tender_id,omitempty"`
	Alg       string  `json:"alg,omitempty"`
	Type      string  `json:"type,omitempty"`
	Procent   float64 `json:"procent"`
	Minimal   float64 `json:"minimal,omitempty"`
	Critical  float64 `json:"critical,omitempty"`
	IsApprove bool    `json:"is_approve,omitempty"`
}
