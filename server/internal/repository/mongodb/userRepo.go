package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"origin-tender-backend/server/internal/domain"
)

func (r repo) GetSiteUserByName(name string) (domain.User, error) {
	var user domain.User

	err := r.userCollection.FindOne(context.Background(),
		bson.D{{"name", name}}).Decode(&user)

	return user, err
}

func (r repo) GetSiteUser(objectId string) (domain.User, error) {
	var user domain.User
	id, _ := primitive.ObjectIDFromHex(objectId)

	err := r.userCollection.FindOne(context.Background(),
		bson.D{{"_id", id}}).Decode(&user)

	return user, err
}
