package logic

import (
	"fmt"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendHello(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.Text == "/start" {
		m :=  "Hello " + update.Message.From.UserName + "!" 
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, m)
		bot.Send(msg)
	}
}

func GiveCurrency(update *tgbotapi.Update, bot *tgbotapi.BotAPI)  {
	if update.Message.Text != "/start" {
		mb, sb, err := GetData(update.Message.Text)

		if err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Currency not found"))
			bot.Send(tgbotapi.NewStickerShare(update.Message.Chat.ID, os.Getenv("404")))
		}else{
		genMsg := fmt.Sprintf("%s\n%s", mb, sb)

		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, genMsg))
		bot.Send(tgbotapi.NewStickerShare(update.Message.Chat.ID, os.Getenv("OK")))
		}
	}
}