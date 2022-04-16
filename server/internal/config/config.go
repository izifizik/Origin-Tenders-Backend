package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

type Config struct {
	App struct {
		Port string `env-default:"8080"`
		Host string `env-default:"0.0.0.0"`
	}
	Database struct {
		Client               *mongo.Client
		TPCollection         *mongo.Collection
		UserCollection       *mongo.Collection
		ProofTokenCollection *mongo.Collection
		TgUserCollection     *mongo.Collection
		TenderCollection     *mongo.Collection
		OrderCollection      *mongo.Collection
		BotCollection        *mongo.Collection
	}
}

var once sync.Once
var instance Config

//NewConfig - create config from env
func NewConfig() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println(err.Error())
		}

		instance.App.Port = "8080"
		instance.App.Host = "0.0.0.0"
		mongoURI := "mongodb://10.10.117.179:27017"
		mongoTokenProofCollection := "token-proof"

		instance.Database.Client, err = mongoConnection(mongoURI)
		instance.Database.TPCollection = instance.Database.Client.Database("Origin-Tenders").Collection(mongoTokenProofCollection)

		instance.Database.BotCollection = instance.Database.Client.Database("Origin-Tenders").Collection("Bot")
		instance.Database.TgUserCollection = instance.Database.Client.Database("Origin-Tenders").Collection("TelegramUsers")
		instance.Database.ProofTokenCollection = instance.Database.Client.Database("Origin-Tenders").Collection("token-proof")
		instance.Database.UserCollection = instance.Database.Client.Database("Origin-Tenders").Collection("Users")
		instance.Database.TenderCollection = instance.Database.Client.Database("Origin-Tenders").Collection("tenders")
		instance.Database.OrderCollection = instance.Database.Client.Database("Origin-Tenders").Collection("orders")
	})
	return instance
}

func mongoConnection(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
