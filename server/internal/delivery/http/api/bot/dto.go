package bot

type TokenDTO struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type TokenProofDTO struct {
	ID    string
	Token string
}
