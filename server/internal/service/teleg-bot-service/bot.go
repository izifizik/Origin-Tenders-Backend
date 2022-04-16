package teleBotService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"origin-tender-backend/server/internal/repository/mongodb"
)

func Run(tgUserRepo mongodb.Repository, tokenProofRepo mongodb.Repository) {
	//513268133
	//tgUser, status, err := tgUserRepo.CreateNewTgUser(513268133, "sas2", "apsdfovoije23")
	//if err != nil {
	//	fmt.Println(status)
	//	fmt.Println(err)
	//}
	//fmt.Println(tgUser.UserId)

	//err := tgUserRepo.UpdateUserStateById("625a0dd3a880a075e99263d8", "fifa2")

	//tgUserRepo.SetUserNameById("625a0dd3a880a075e99263d8", "FUVA")

	//tokenProofRepo.CreateToken("Dimasik", "Chudinov")
	//
	//status, err := tokenProofRepo.ApproveProofToken("Dimasik", "Chudinov")
	//
	//status2, err := tokenProofRepo.ApproveProofToken("Dimasik", "Chudinov2")

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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "no text")

			user, err := tgUserRepo.GetUserByTgId(update.SentFrom().ID)
			if err != nil {
				fmt.Println("get user db error")
			}

			if user.UserId == 0 {
				msg.Text = "введите ваш username с сайта"

				tgUser, _, _ := tgUserRepo.CreateNewTgUser(update.SentFrom().ID, "", "")

				tgUserRepo.UpdateUserStateById(tgUser.Id, "entering_token")

				bot.Send(msg)
				continue
			}

			switch user.State {
			case "entering_name":
				tgUser, _, _ := tgUserRepo.GetTgUser(update.Message.Text)

				if tgUser.UserId == 0 {
					tgUserRepo.SetUserNameById(user.Id, update.Message.Text)
					msg.Text = "введите токен"
					tgUserRepo.UpdateUserStateById(user.Id, "entering_name")
				} else {
					msg.Text = "такой пользователь уже зарегистрирован в телеграм боте, введите другое имя"

				}

			case "entering_token":
				token := update.Message.Text

				status, err := tokenProofRepo.ApproveProofToken(user.Name, token)
				if err != nil {
					fmt.Println(err)
				}

				if status == "invalid token" {
					msg.Text = "incorrect token, try again"
					continue
				}

				tgUserRepo.SetSiteId(status, user.Id)
				msg.Text = "вы успешно связались с сервисом, ждите новостей!"
			}

			bot.Send(msg)
		}
	}
}
