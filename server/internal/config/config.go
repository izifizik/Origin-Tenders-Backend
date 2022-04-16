package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	App struct {
		Port string `env-default:"8080"`
		Host string `env-default:"localhost"`
	}
	Database struct {
		Client *mongo.Client
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

		instance.App.Port = os.Getenv("PORT")
		instance.App.Host = os.Getenv("HOST")
		mongoURI := os.Getenv("MONGO_URI")

		instance.Database.Client, err = mongoConnection(mongoURI)
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
