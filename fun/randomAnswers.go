package fun

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRandomAnswer() string {

	url := "https://yesno.wtf/api"

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
	resp = fmt.Sprintf("%v", result["image"])
	return resp
}
