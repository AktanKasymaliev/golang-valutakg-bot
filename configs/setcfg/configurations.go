package setcfg

import "os"


func SetCfg() {
	os.Setenv("TG_TOKEN", "2033268440:AAEywcWdOTsXUjjF5oWe6yvrxMwavmYKDzQ")
	os.Setenv("BANK_API", "https://banks.kg/api/v1/rates/daily/usd")
}