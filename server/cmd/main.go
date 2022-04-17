package main

import (
	"fmt"
	"log"
	"origin-tender-backend/server/internal/app"
)

func main() {

	fmt.Println("!!!")

	//go websocket.Run()

	err := app.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
