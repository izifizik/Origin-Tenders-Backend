package domain

type TelegramUser struct {
	Id     string `bson:"_id"`
	SiteId string
	UserId int64
	Name   string // unique site name
	State  string
}
