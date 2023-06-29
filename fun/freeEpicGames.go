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

//TODO Переделать метод с апи rapidapi, на метод о эпиков. Работает просто GET, без авторизации
/* curl 'https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=RU&allowCountries=RU' \
-H 'authority: store-site-backend-static-ipv4.ak.epicgames.com' \
-H 'accept: application/json, text/plain, *\/*' \
-H 'accept-language: ru-RU,ru;q=0.7' \
-H 'origin: https://store.epicgames.com' \
-H 'referer: https://store.epicgames.com/' \
-H 'sec-ch-ua: "Not.A/Brand";v="8", "Chromium";v="114", "Brave";v="114"' \
-H 'sec-ch-ua-mobile: ?0' \
-H 'sec-ch-ua-platform: "Windows"' \
-H 'sec-fetch-dest: empty' \
-H 'sec-fetch-mode: cors' \
-H 'sec-fetch-site: same-site' \
-H 'sec-gpc: 1' \
-H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36' \
-H 'x-requested-with: XMLHttpRequest' \
--compressed
*/
