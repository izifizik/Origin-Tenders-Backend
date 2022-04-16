package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"origin-tender-backend/server/internal/repository/mongodb"
)

func HandleMessages(update tgbotapi.Update, bot *tgbotapi.BotAPI, repo mongodb.Repository) {

	//update.Message.Chat.Type
	//err := actions.SendAcceptParticipationInTender(update.SentFrom().ID, " sas", 223)
	// group

	if update.Message.Chat.Type == "group" {
		handleGroupMessages(update, bot, repo)
	} else {
		handlePrivateMessages(update, bot, repo)
	}

}

func handlePrivateMessages(update tgbotapi.Update, bot *tgbotapi.BotAPI, repo mongodb.Repository) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "no text")

	user, err := repo.GetUserByTgId(update.SentFrom().ID)
	if err != nil {
		fmt.Println("get user db error")
	}

	if user.UserId == 0 {
		userName := update.SentFrom().UserName
		fmt.Println(userName)
		msg.Text = "👋 Это бот для автоматизации получения уведомлений 👋\n с сайта zakupku.mos.ru, пожалуйста, \n Введите  введите ваш токен \n с сайта для получения доступа к функциям бота 🌚"

		tgUser, _, _ := repo.CreateNewTgUser(update.SentFrom().ID, "", "")

		repo.UpdateUserStateById(tgUser.Id, "entering_token")

		bot.Send(msg)
		return
	}

	switch user.State {
	case "entering_name":
		tgUser, _, _ := repo.GetTgUser(update.Message.Text)

		if tgUser.UserId == 0 {
			repo.SetUserNameById(user.Id, update.Message.Text)
			msg.Text = "введите токен"
			repo.UpdateUserStateById(user.Id, "entering_name")
		} else {
			msg.Text = "такой пользователь уже зарегистрирован в телеграм боте, введите другое имя"

		}

	case "":
		token := update.Message.Text

		status, err := repo.ApproveProofToken(update.SentFrom().UserName, token)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(update.SentFrom().UserName + " !!!!!!!!!!!!!!!!!!!")

		if status == "invalid token" {
			msg.Text = "вы успешно связались с сервисом, ждите новых уведомлений!!!"
			bot.Send(msg)
			return
		}

		repo.SetSiteId(status, user.Id)
		msg.Text = "вы успешно связались с сервисом, ждите новых уведомлений!"

		repo.UpdateUserStateById(user.Id, "main")
	}

	bot.Send(msg)
}

func handleGroupMessages(update tgbotapi.Update, bot *tgbotapi.BotAPI, repo mongodb.Repository) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "no text")

	groupUser, err := repo.GetUserByTgId(update.Message.Chat.ID)
	if err != nil {
		fmt.Println(err)
	}

	if groupUser.UserId == 0 {
		repo.CreateNewTgUser(update.Message.Chat.ID, update.Message.Chat.UserName, "")
		msg.Text = "дратути это бот для спама, пожалуйста скиньте нам констринг на монгу"
	} else {
		msg.Text = "Отлично, вы привязали спам к этому чату!"
	}

	bot.Send(msg)
}
