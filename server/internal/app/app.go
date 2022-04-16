package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"origin-tender-backend/server/internal/config"
	"origin-tender-backend/server/internal/delivery/http/api/bot"
	"origin-tender-backend/server/internal/repository/mongodb"
	botService "origin-tender-backend/server/internal/service/bot-service"
	teleBotService "origin-tender-backend/server/internal/service/teleg-bot-service"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func Run() error {

	cfg := config.NewConfig()
	//go websocket.Run()

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        router,
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println(cfg.App.Port, cfg.App.Host)
	repo := mongodb.NewRepo(cfg.Database.Client, cfg.Database.TPCollection, cfg.Database.UserCollection,
		cfg.Database.ProofTokenCollection, cfg.Database.TgUserCollection, cfg.Database.TenderCollection,
		cfg.Database.OrderCollection, cfg.Database.BotCollection)

	service := botService.NewBotService(repo)

	//service.CreateTgToken("Necromant24", "dflk", )

	//err := service.CreateSiteUser(domain.User{Name: "Dimasik)"})
	//user, err := service.GetSiteUserByName("Dimasik)")
	//
	//fmt.Println(user.ID, err)

	go teleBotService.Run(repo)

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
