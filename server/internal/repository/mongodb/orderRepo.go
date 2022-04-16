package mongodb

import (
	"context"
	"origin-tender-backend/server/internal/domain"
)

func (r repo) CreateOrder(order domain.Order) error {
	//_, err := r.proofTokenCollection.InsertOne(context.Background(), bson.D{
	//	{"timeStamp", time.Now()},
	//	{"userId", order.UserId},
	//	{"userName", order.UserName},
	//	{"TenderId", order.TenderId},
	//	{"TenderName", order.TenderName},
	//	{"Price", order.Price},
	//})

	_, err := r.ordersCollection.InsertOne(context.Background(), order)

	// TODO: do stuff, call bot events

	return err
}
