package fun

import (
	"io"
	"math/rand"
	"net/http"
	"time"
)

func GenerateInsult() string {

	url := "https://evilinsult.com/generate_insult.php?lang=ru"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	persons := []string{"Пабло, ", "Хасан, ", "Гайдар, ", "Михан, ", "Эдик, ", "Женя, ", "Сева, "}

	rand.Seed(time.Now().UnixNano())

	var varPersons string
	for i := 0; i < 3; i++ {
		randomIndex := rand.Intn(len(persons))
		varPersons = persons[randomIndex]

	}
	forReturn := varPersons + string(body)
	return forReturn

}
