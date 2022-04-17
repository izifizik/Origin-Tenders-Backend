package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
)

func (r repo) CreateOrder(order domain.Order) error {
	_, err := r.ordersCollection.InsertOne(context.Background(), order)
	if err != nil {
		return err
	}
	userId, err := primitive.ObjectIDFromHex(order.TenderId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": userId}

	tender, err := r.GetTenderByID(order.TenderId)

	tender.CurrentPrice = order.Price

	err = r.UpdateTender(filter, tender)

	return err
}

	return nil
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
	err = c.All(context.Background(), orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
