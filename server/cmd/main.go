package main

import (
	"fmt"
	"log"
	"origin-tender-backend/server/internal/app"
	"origin-tender-backend/websocket"
)

func main() {

	fmt.Println("!!!")

	websocket.Run()

	err := app.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
