package fun

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRandomGif() string {
	url := "https://api.giphy.com/v1/gifs/random?api_key=vIeNhqLVATaSiUyWbtjYQBYwEJ54S4zC&tag=&rating=g"

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
	resp = fmt.Sprintf("%v", result["data"].(map[string]interface{})["bitly_gif_url"])
	return resp
}
