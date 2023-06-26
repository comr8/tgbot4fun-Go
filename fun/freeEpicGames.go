package fun

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetFreeGames() string {

	url := "https://free-epic-games.p.rapidapi.com/free"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "35f0bc0607msh248e5120529c4e6p1e7331jsn42dc9b08a777")
	req.Header.Add("X-RapidAPI-Host", "free-epic-games.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	var resp string
	resp = fmt.Sprintf("%v", result["freeGames"].(map[string]interface{})["current"].([]interface{})[0].(map[string]interface{})["title"])
	resp = "Бесплатная игра в EpicGames: " + resp
	return resp
}
