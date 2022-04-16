package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	userId, err := primitive.ObjectIDFromHex(order.TenderId)
	filter := bson.M{"_id": userId}

	_, err = r.tendersCollection.UpdateOne(context.Background(), filter,
		bson.D{{"current_price", order.Price}})

	// TODO: do stuff, call bot events

	return err
}
