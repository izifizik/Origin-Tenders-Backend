package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

var TgBot *tgbotapi.BotAPI

func SendAcceptParticipationInTender(tgUserId int64, tenderName string, price float32) error {

	var yesNo = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("yes", "1"),
			tgbotapi.NewInlineKeyboardButtonData("no", "0"),
		),
	)

	// tgUserId - chatId?
	var priceStr = strconv.Itoa(int(price))
	msg := tgbotapi.NewMessage(tgUserId, "название тендера - "+tenderName+"\nцена: "+priceStr)
	msg.ReplyMarkup = yesNo
	_, err := TgBot.Send(msg)

	return err
}
