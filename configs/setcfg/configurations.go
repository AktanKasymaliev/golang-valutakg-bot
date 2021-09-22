package setcfg

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func parseCfg(section string, key string) (string) {
	cfg, err := ini.Load("settings.ini")
	if err != nil {
		log.Fatal(err)
    }
	return cfg.Section(section).Key(key).String()
}

func SetCfg() {
	os.Setenv("TG_TOKEN", parseCfg("BOT", "TG_TOKEN"))
	os.Setenv("OK", parseCfg("BOT", "OK"))
	os.Setenv("404", parseCfg("BOT", "404"))

	os.Setenv("BANK_API", parseCfg("API", "BANK_API"))
}