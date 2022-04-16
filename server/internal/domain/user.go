package domain

type User struct {
	ID   string `bson:"_id"`
	Name string
	Notifications
	Filters        []string
	TendersHistory []Tender
}

type Notifications struct {
	TgID string
}
