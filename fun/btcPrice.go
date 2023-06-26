package fun

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBtcPrice() string {
	url := "https://api.coinranking.com/v2/coin/Qwsogvtv82FCd/price"
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
	resp = "Биток стоит: " + fmt.Sprintf("%v", result["data"].(map[string]interface{})["price"]) + " USD"

	return resp
}
