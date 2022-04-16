package domain

type User struct {
	ID             string `bson:"_id"`
	Name           string
	Filters        []string
	TendersHistory []Tender
}
