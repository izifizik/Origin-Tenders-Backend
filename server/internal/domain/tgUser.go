package domain

type TelegramUser struct {
	UserId int64
	Name   string // unique site name
	State  string
}
