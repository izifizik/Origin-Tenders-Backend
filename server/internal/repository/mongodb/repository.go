package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"origin-tender-backend/server/internal/domain"
)

type repo struct {
	client               *mongo.Client
	tpCollection         *mongo.Collection
	userCollection       *mongo.Collection
	proofTokenCollection *mongo.Collection
	tgUserCollection     *mongo.Collection
	tendersCollection    *mongo.Collection
	ordersCollection     *mongo.Collection
	botCollection        *mongo.Collection
}

func NewRepo(client *mongo.Client, tpCollection, userCollection, proofTokenCollection,
	tgUserCollection, tenderCollection,
	ordersCollection, botCollection *mongo.Collection) Repository {
	return &repo{
		client, tpCollection, userCollection,
		proofTokenCollection, tgUserCollection,
		tenderCollection, ordersCollection, botCollection,
	}
}

func (r repo) CreateSiteUser(user domain.User) error {
	_, err := r.userCollection.InsertOne(context.Background(), bson.D{
		{"name", user.Name},
	})

	return err
}

func (r repo) SaveToken(ID string, token string) error {
	_, err := r.tpCollection.InsertOne(context.Background(), token)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) GetUserByTgId(id int64) (domain.TelegramUser, error) {
	var tgUser domain.TelegramUser

	err := r.tgUserCollection.FindOne(context.Background(), bson.D{
		{"userId", id},
	}).Decode(&tgUser)

	if err != nil {
		return tgUser, err
	}

	return tgUser, nil
}

func (r repo) GetTgUser(name string) (domain.TelegramUser, string, error) {
	var tgUser domain.TelegramUser

	err := r.tgUserCollection.FindOne(context.Background(), bson.D{
		{"name", name},
	}).Decode(&tgUser)
	if err != nil {
		return tgUser, "db error", err
	}

	return tgUser, "ok", nil
}

// return (tgUser,status, error)
func (r repo) CreateNewTgUser(id int64, name string, token string) (domain.TelegramUser, string, error) {

	var tgUser domain.TelegramUser

	err := r.tgUserCollection.FindOne(context.Background(), bson.D{
		{"userId", id},
	}).Decode(&tgUser)
	if err != nil {
		fmt.Println(tgUser, "db error: "+err.Error())
		return tgUser, "db error: " + err.Error(), err
	}

	_, err = r.tgUserCollection.InsertOne(context.Background(), bson.D{
		{"userId", id},
		{"name", name},
		{"token", token},
	})

	if err != nil {
		return tgUser, "db error: " + err.Error(), err
	}

	return tgUser, "success", nil
}

func (r repo) SetUserNameById(idStr string, name string) error {
	var tgUser domain.TelegramUser

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", name}}}}

	_, err = r.tgUserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(tgUser, "db error: "+err.Error())
	}

	return err
}

func (r repo) SetSiteId(siteId string, idStr string) error {
	var tgUser domain.TelegramUser

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"siteId", siteId}}}}

	_, err = r.tgUserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(tgUser, "db error: "+err.Error())
	}

	return err
}

func (r repo) UpdateUserStateById(idStr string, state string) error {
	var tgUser domain.TelegramUser

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"state", state}}}}

	_, err = r.tgUserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(tgUser, "db error: ", err.Error())
	}

	return err
}

func (r repo) CreateToken(name string, token string, siteId primitive.ObjectID) error {
	var proofedToken domain.ProofToken

	err := r.proofTokenCollection.FindOne(context.Background(), bson.D{
		{"name", name},
	}).Decode(&proofedToken)
	if err != nil {
		return err
	}
	if proofedToken.Id != "" {
		r.proofTokenCollection.DeleteOne(context.Background(), bson.M{"name": name})
	}

	_, err = r.proofTokenCollection.InsertOne(context.Background(), bson.D{
		{"name", name},
		{"token", token},
		{"siteId", siteId},
	})

	return err
}

func (r repo) ApproveProofToken(name string, token string) (string, error) {
	var proofToken domain.ProofToken
	status := "ok"

	err := r.proofTokenCollection.FindOne(context.Background(), bson.D{
		{"name", name},
	}).Decode(&proofToken)

	if proofToken.Token != token {
		status = "invalid token"
	} else {
		status = proofToken.SiteId
	}

	return status, err
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
