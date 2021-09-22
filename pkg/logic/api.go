package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)


func GetData(currency string) (string, string, error) {
	d, err := makeRequest(os.Getenv("BANK_API") + currency)

	if err != nil {
		return "", "", err
	}

	mb, sb, err := parseData(d)

	if err != nil {
		return "", "", err
	}

	return mb, sb, nil
}

func parseData(bodyBytes []byte) (string, string, error) {

	var data map[string][][]float64
	err := json.Unmarshal(bodyBytes, &data)

	if err != nil {
		return "", "", err
	}

	sell := data["sell"]
	buy := data["buy"]
	currentMonth := time.Unix(int64(buy[len(buy)-3][0]/1000), 0).Month()
	currentDate := fmt.Sprintf("%d:%02d - %s, %d", time.Now().Hour(), time.Now().Minute(), currentMonth, time.Now().Day(),)

	mb := fmt.Sprintf("(%s) %s %v som", currentDate, "Purchase rate: ", buy[len(buy)-1][1])
	sb := fmt.Sprintf("(%s) %s %v som", currentDate, "Sale rate: ", sell[len(sell)-1][1])

	return mb, sb, nil
}

func makeRequest(url string) ([]byte, error) {

	r, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if r.StatusCode == http.StatusOK {
		defer r.Body.Close()
		bodyBytes, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
		}

		return bodyBytes, nil
	}

	return nil, err
}