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

func (r repo) GetOrderById(objectId string) (domain.Order, error) {
	var order domain.Order
	userId, err := primitive.ObjectIDFromHex(objectId)

	filter := bson.M{"_id": userId}
	err = r.ordersCollection.FindOne(context.Background(), filter).Decode(&order)

	return order, err
}

func (r repo) GetTenderOrders(tenderId string) ([]domain.Order, error) {
	var orders []domain.Order
	c, err := r.ordersCollection.Find(context.Background(), bson.D{{"tenderId", tenderId}})

	err = c.All(context.Background(), orders)

	return orders, err
}
