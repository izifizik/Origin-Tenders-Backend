package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	client *mongo.Client
}

func NewRepo(client *mongo.Client) Repository {
	return &repo{client}
}

func (r repo) SaveToken(ctx context.Context, ID string, file string) {
	_, err := fileCollection.InsertOne(context.Background(), file)
	if err != nil {
		return err
	}
	return nil
}
