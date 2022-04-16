package domain

type ProofToken struct {
	Id     string `bson:"_id"`
	Name   string
	Token  string
	SiteId string
}
