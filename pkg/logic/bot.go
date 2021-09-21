package logic

import (

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendHello(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.Text == "/start" {
		m :=  "Hello " + update.Message.From.UserName + "!" 
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, m)
		bot.Send(msg)
	}
}