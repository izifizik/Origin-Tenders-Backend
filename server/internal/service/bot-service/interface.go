package botService

type BotService interface {
	GenerateToken(ID string) string
	ProofToken(ID string, token string) error
}
