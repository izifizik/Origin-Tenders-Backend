package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	client       *mongo.Client
	tpCollection *mongo.Collection
}

func NewRepo(client *mongo.Client, tpCollection *mongo.Collection) Repository {
	return &repo{client, tpCollection}
}

func (r repo) SaveToken(ID string, token string) error {
	_, err := r.tpCollection.InsertOne(context.Background(), token)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) ProofToken(ID string, token string) error {
	defer r.DeleteToken(ID)
	var dto struct {
		ID    string
		Token string
	}
	filter := bson.M{"ID": ID}

	err := r.tpCollection.FindOne(context.Background(), filter).Decode(&dto)
	if err != nil {
		return fmt.Errorf("error with find by id")
	}

	if dto.Token != token {
		return fmt.Errorf("error proof tokens is not equal")
	}
	return nil
}

func (r repo) DeleteToken(ID string) {
	filter := bson.M{"ID": ID}

	_, err := r.tpCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("error with delete token by ID: " + err.Error())
	}
}
