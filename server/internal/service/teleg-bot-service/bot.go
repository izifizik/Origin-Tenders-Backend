package teleBotService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {
	bot, err := tgbotapi.NewBotAPI("apiToken")
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
