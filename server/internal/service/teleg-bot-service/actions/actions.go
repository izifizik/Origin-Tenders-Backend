package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"origin-tender-backend/server/internal/domain"
	"strconv"
)

var TgBot *tgbotapi.BotAPI

func SendAcceptParticipationInTender(tgUserId int64, tender domain.Tender) error {

	var yesNo = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("подтвердить", "Участие подтверждено"),
			tgbotapi.NewInlineKeyboardButtonData("отклонить", "Участие отклонено"),
		),
	)

	// tgUserId - chatId?
	var priceStr = strconv.Itoa(int(tender.StartPrice))
	msg := tgbotapi.NewMessage(tgUserId, "название тендера - "+tender.Name+"\nцена: "+priceStr+" Р\nшаг (в процентах): "+strconv.FormatFloat(float64(tender.StepPercent), 'f', 6, 64)+" %\n описание: "+tender.ShortDescription)
	msg.ReplyMarkup = yesNo
	_, err := TgBot.Send(msg)

	return err
}
