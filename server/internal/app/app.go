package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"origin-tender-backend/server/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
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
