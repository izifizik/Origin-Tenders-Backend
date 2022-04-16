package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"origin-tender-backend/server/internal/config"

	teleBotService "origin-tender-backend/server/internal/service/teleg-bot-service"

	"origin-tender-backend/server/internal/delivery/http/api/bot"
	"origin-tender-backend/server/internal/repository/mongodb"
	botService "origin-tender-backend/server/internal/service/bot-service"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://10.10.117.179:27017"))
	if err != nil {
		fmt.Println("not found database")
		panic(err)
	}

	collection := client.Database("Origin-Tenders").Collection("TelegramUsers")
	if collection == nil {
		panic("collection is nill")
	}

	tokenCollection := client.Database("Origin-Tenders").Collection("token-proof")
	if tokenCollection == nil {
		panic("collection is nill")
	}

	var tgUserRepo mongodb.Repository = mongodb.NewRepo(client, collection)
	var tokenProofRepo mongodb.Repository = mongodb.NewRepo(client, tokenCollection)

	teleBotService.Run(tgUserRepo, tokenProofRepo)

	cfg := config.NewConfig()

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	server := &http.Server{
		Addr:           cfg.App.Host + ":" + cfg.App.Port,
		Handler:        router,
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: 1 << 20,
	}

	repo := mongodb.NewRepo(cfg.Database.Client, cfg.Database.TPCollection)

	service := botService.NewBotService(repo)

	handler := bot.NewBotHandler(service)

	handler.Register(router)

	go gracefulShutdown([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}, server)

	return server.ListenAndServe()
}

func gracefulShutdown(signals []os.Signal, closeItems ...io.Closer) {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, signals...)
	sig := <-sign

	log.Printf("Caught signal %s. Shutting down...", sig)
	for _, closer := range closeItems {
		err := closer.Close()
		if err != nil {
			fmt.Printf("failed to close %v: %v", closer, err)
		}
	}
}
