package domain

import "time"

type Order struct {
	TimeStamp  time.Time `json:"timeStamp" bson:"timeStamp"`
	UserId     string    `json:"userId" bson:"userId"`
	UserName   string    `json:"userName" bson:"userName"`
	TenderId   string    `json:"tenderId" bson:"tenderId"`
	TenderName string    `json:"tenderName" bson:"tenderName"`
	Price      float32   `json:"price" bson:"price"`
}
