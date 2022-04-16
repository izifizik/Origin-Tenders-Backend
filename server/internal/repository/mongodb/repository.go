package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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

//func (r repo) ProofToken(ctx context.Context, ID string, file string) (bool, error) {
//	filter := bson.M{"tg": name}
//
//	err := clubsCollection.FindOne(context.Background(), filter).Decode(&club)
//	if err != nil {
//		fmt.Println(err)
//		return model.Club{}
//	}
//	return club
//}
