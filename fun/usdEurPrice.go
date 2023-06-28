package fun

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUsdRate() string {
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result map[string]interface{}
	err := json.Unmarshal(body, &result)

	if err != nil {
		panic(err)
	}

	var resp string
	resp = "Доллар: " + fmt.Sprintf("%v", result["Valute"].(map[string]interface{})["USD"].(map[string]interface{})["Value"]) + " рублей (по ЦБ)"

	return resp

}

func GetEuRate() string {
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result map[string]interface{}
	err := json.Unmarshal(body, &result)

	if err != nil {
		panic(err)
	}

	var resp string
	resp = "Евро: " + fmt.Sprintf("%v", result["Valute"].(map[string]interface{})["EUR"].(map[string]interface{})["Value"]) + " рублей (по ЦБ)"
	return resp
}

//TODO Добавить кнопки для получения курса валют и битка
