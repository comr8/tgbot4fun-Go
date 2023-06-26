package fun

import (
	"math/rand"
	"time"
)

func RandomEventsGenerator4Weekend() string {
	events := []string{"Поплакать и поржать над мемами.", "Выпить.", "Ахуеть.",
		"Следить за ЧВК Вагнер на Google Maps.", "Покататься на танке.", "Заказать билеты в Стамбул за 90к руб.",
		"Катануть в мордаху.", "Учиться и работать над собой, потратить день с пользой.", "Встретиться с семьей.",
		"Начать читать \"Капитал\" К.Маркса"}

	rand.Seed(time.Now().UnixNano())
	var eventVar string
	for i := 0; i < 3; i++ {
		randomIndex := rand.Intn(len(events))
		eventVar = events[randomIndex]
	}

	return eventVar
}
