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
}

func NewRepo(client *mongo.Client, tpCollection *mongo.Collection,
	userCollection *mongo.Collection, proofTokenCollection *mongo.Collection,
	tgUserCollection *mongo.Collection) Repository {
	return &repo{
		client, tpCollection, userCollection,
		proofTokenCollection, tgUserCollection,
	}
}

func (r repo) CreateSiteUser(user domain.User) error {
	_, err := r.userCollection.InsertOne(context.Background(), bson.D{
		{"name", user.Name},
	})

	return err
}

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
		return tgUser, nil
	}

	return tgUser, nil
}

func (r repo) GetTgUser(name string) (domain.TelegramUser, string, error) {
	var tgUser domain.TelegramUser

	err := r.tgUserCollection.FindOne(context.Background(), bson.D{
		{"name", name},
	}).Decode(&tgUser)

	if err != nil {
		return tgUser, "db error", nil
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
		fmt.Println(tgUser, "db error", nil)
	}

	_, err2 := r.tgUserCollection.InsertOne(context.Background(), bson.D{
		{"userId", id},
		{"name", name},
		{"token", token},
	})

	if err2 != nil {
		return tgUser, "db error", nil
	}

	return tgUser, "success", nil
}

func (r repo) SetUserNameById(idStr string, name string) error {
	var tgUser domain.TelegramUser

	id, _ := primitive.ObjectIDFromHex(idStr)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", name}}}}

	_, err := r.tgUserCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		fmt.Println(tgUser, "db error", nil)
	}

	return err
}

func (r repo) SetSiteId(siteId string, idStr string) error {
	var tgUser domain.TelegramUser

	id, _ := primitive.ObjectIDFromHex(idStr)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"siteId", siteId}}}}

	_, err := r.tgUserCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		fmt.Println(tgUser, "db error", nil)
	}

	return err
}

func (r repo) UpdateUserStateById(idStr string, state string) error {
	var tgUser domain.TelegramUser

	id, _ := primitive.ObjectIDFromHex(idStr)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"state", state}}}}

	_, err := r.tgUserCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		fmt.Println(tgUser, "db error", nil)
	}

	return err
}

//func (r repo) UpdateUserState(id int64, state string) error {
//	r.tpCollection.UpdateByID(context.Background(), bson.D{
//		{""},
//	})
//}

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

func (r repo) CreateToken(name string, token string, siteId string) error {

	_, err := r.proofTokenCollection.InsertOne(context.Background(), bson.D{
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
