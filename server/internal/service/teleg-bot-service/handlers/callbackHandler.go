package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var yes = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("подтвердить", "Участие подтверждено"),
	),
)

var no = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("отклонить", "Участие отклонено"),
	),
)

func HandleCallBack(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery.Data != "" {

		//var action domain.TgAction
		//json.Unmarshal([]byte(update.CallbackQuery.Data), &action)

		switch update.CallbackQuery.Message.Text {
		case "подтвердить":
			update.Message.ReplyMarkup = &yes
		case "отклонить":
			update.Message.ReplyMarkup = &no

		}

		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
		msg.ReplyToMessageID = update.CallbackQuery.Message.MessageID

		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		if _, err := bot.Request(callback); err != nil {
			panic(err)
		}

		// And finally, send a message containing the data received.
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}

	}
}
