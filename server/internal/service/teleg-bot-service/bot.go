package teleBotService

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"origin-tender-backend/server/internal/repository/mongodb"
	"origin-tender-backend/server/internal/service/teleg-bot-service/actions"
	"origin-tender-backend/server/internal/service/teleg-bot-service/handlers"
)

func Run(repo mongodb.Repository) {
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

	actions.TgBot = bot

	for update := range updates {

		if update.CallbackQuery != nil {
			handlers.HandleCallBack(update, bot)
		}

		if update.Message != nil {
			handlers.HandleMessages(update, bot, repo)
		}
	}
}
