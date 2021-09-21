package main

import (

	"tgbot/configs/setcfg"
	"tgbot/internal/app"
	"tgbot/pkg/logic"

)

func main() {

	setcfg.SetCfg()
	bot, updates := app.InitBot()
	
	for update := range updates {

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		logic.SendHello(&update, bot)

	}
}
