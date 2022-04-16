package teleBotService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"origin-tender-backend/server/internal/repository/mongodb"
)

func Run(repo mongodb.Repository) {
	//513268133
	tgUser, status, err := repo.CreateNewTgUser(513268133, "sas", "apsdfovoije23")
	if err != nil {
		fmt.Println(status)
		fmt.Println(err)
	}

	fmt.Println(tgUser.UserId)

	bot, err := tgbotapi.NewBotAPI("493403387:AAGsQne6Pj0NTTQPwYo95gZ23jx5A2t59pw")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	fmt.Println("start...")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					msg.Text = "введите ваш username с сайта"

				}
			}

			bot.Send(msg)
		}
	}
}
