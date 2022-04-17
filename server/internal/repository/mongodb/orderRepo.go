package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
	"time"
)

func (r repo) CreateOrder(order domain.Order) error {
	order.TimeStamp = time.Now()
	res, err := r.ordersCollection.InsertOne(context.Background(), order)
	if err != nil {
		return err
	}

	fmt.Println(res.InsertedID)

	//id := string([]byte(res.InsertedID))
	//fmt.Println(id)

	return err
}

func (r repo) GetOrderById(objectId string) (domain.Order, error) {
	var order domain.Order
	userId, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return domain.Order{}, err
	}
	filter := bson.M{"_id": userId}
	err = r.ordersCollection.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (r repo) GetTenderOrders(tenderId string) ([]domain.Order, error) {
	var orders []domain.Order
	c, err := r.ordersCollection.Find(context.Background(), bson.D{{"tenderId", tenderId}})
	if err != nil {
		return nil, err
	}

	err = c.All(context.Background(), &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
