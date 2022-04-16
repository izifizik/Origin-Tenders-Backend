package botService

type BotService interface {
	GenerateToken(ID string) string
}
