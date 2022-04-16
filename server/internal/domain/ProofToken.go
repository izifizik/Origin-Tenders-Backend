package domain

type ProofToken struct {
	Id     string `bson:"_id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Token  string `bson:"token" bson:"token"`
	SiteId string `bson:"siteId" json:"siteId"`
}
