package main

import (
	"log"
	"origin-tender-backend/server/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
