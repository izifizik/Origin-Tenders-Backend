package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"origin-tender-backend/server/internal/domain"
)

type repo struct {
	client       *mongo.Client
	tpCollection *mongo.Collection
}

func NewRepo(client *mongo.Client, tpCollection *mongo.Collection) Repository {
	return &repo{client, tpCollection}
}

func (r repo) SaveToken(ctx context.Context, ID string, file string) error {
	_, err := r.tpCollection.InsertOne(context.Background(), file)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) GetTgUser(name string) (domain.TelegramUser, string, error) {
	var tgUser domain.TelegramUser

	err := r.tpCollection.FindOne(context.Background(), bson.D{
		{"name", name},
	})

	if err != nil {
		return tgUser, "db error", nil
	}

	return tgUser, "ok", nil
}

// return (tgUser,status, error)
func (r repo) CreateNewTgUser(id int64, name string, token string) (domain.TelegramUser, string, error) {

	var tgUser domain.TelegramUser

	err := r.tpCollection.FindOne(context.Background(), bson.D{
		{"id", id},
	}).Decode(&tgUser)

	if err != nil {
		return tgUser, "db error", nil
	}

	_, err2 := r.tpCollection.InsertOne(context.Background(), bson.D{
		{"id", id},
		{"name", name},
		{"token", token},
	})

	if err2 != nil {
		return tgUser, "db error", nil
	}

	return tgUser, "success", nil
}
