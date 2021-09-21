package app

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel)  {
	
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	updates, err := PoolUpdate(*bot)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot, updates

}

func PoolUpdate(bot tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		return nil, err
	}

	return updates, nil
}