package domain

import "time"

type Order struct {
	TimeStamp  time.Time
	UserId     string
	UserName   string
	TenderId   string
	TenderName string
	Price      float32
}
