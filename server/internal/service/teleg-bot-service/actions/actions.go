package actions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"origin-tender-backend/server/internal/domain"
	"strconv"
)

var TgBot *tgbotapi.BotAPI

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func NotificateTenderChange(tgUserId int64, tender domain.Tender) error {

	// tgUserId - chatId?
	var priceStr = strconv.Itoa(int(tender.StartPrice))
	msg := tgbotapi.NewMessage(tgUserId, "тендер обновлен! \n название тендера - "+tender.Name+"\nцена: "+priceStr+" Р\nшаг (в процентах): "+FloatToString(tender.StepPercent)+" %\n описание: "+tender.ShortDescription)

	_, err := TgBot.Send(msg)

	return err

}

func SendAcceptParticipationInTender(tgUserId int64, tender domain.Tender) error {

	var yesNo = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("подтвердить", "Участие подтверждено"),
			tgbotapi.NewInlineKeyboardButtonData("отклонить", "Участие отклонено"),
		),
	)

	// tgUserId - chatId?
	var priceStr = strconv.Itoa(int(tender.StartPrice))
	msg := tgbotapi.NewMessage(tgUserId, "название тендера - "+tender.Name+"\nцена: "+priceStr+" Р\nшаг (в процентах): "+FloatToString(tender.StepPercent)+" %\n описание: "+tender.ShortDescription)
	msg.ReplyMarkup = yesNo
	_, err := TgBot.Send(msg)

	return err
}
