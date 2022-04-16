package domain

import "time"

type Order struct {
	TimeStamp  time.Time `json:"time_stamp" bson:"timeStamp"`
	UserId     string    `json:"user_id" bson:"userId"`
	UserName   string    `json:"user_name" bson:"userName"`
	TenderId   string    `json:"tender_id" bson:"tenderId"`
	TenderName string    `json:"tender_name" bson:"tenderName"`
	Price      float32   `json:"price" bson:"price"`
}
