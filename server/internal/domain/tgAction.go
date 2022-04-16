package domain

type TgAction struct {
	Type  string `json:"type"`
	Check bool   `json:"check"`
	Data  string `json:"data"`
}
