package main

import (
	"fmt"
	"log"
	"origin-tender-backend/server/internal/app"
	teleBotService "origin-tender-backend/server/internal/service/teleg-bot-service"
)

func main() {

	teleBotService.Run()

	fmt.Println("!!!")

	err := app.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
