package telega

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	fun "../fun"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMessage(chatID int64, message string, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func TelegramBot() {
	bot, err := tgbotapi.NewBotAPI(TgBotToken) /*const TgBotToken из ./telega/conf.go
	dGhlX21vc3RfdG9sZXJhbnRfYm90 --base64	= the_most_tolerant_bot
		name: dGhlX21vc3RfdG9sZXJhbnRfYm90_bot
			t.me/dGhlX21vc3RfdG9sZXJhbnRfYm90_bot */
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		var resp string
		if strings.Contains(msg.Text, "/insultsmb") {
			names := strings.Split(msg.Text, " ")
			var namesForReturn string
			if len(names) >= 2 {
				namesForReturn = names[1]
			}
			resp = namesForReturn + fun.GenerateInsult()
		} else if strings.Contains(msg.Text, "/freegameslist") {
			resp = fun.GetFreeGames()
		} else if strings.Contains(msg.Text, "/randomanswer") {
			resp = fun.GetRandomAnswer()
		} else if strings.Contains(msg.Text, "/randomgif") {
			resp = fun.GetRandomGif()
		} else if strings.Contains(msg.Text, "/btcprice") {
			resp = fun.GetBtcPrice()
		} else if strings.Contains(msg.Text, "/usdrate") {
			resp = fun.GetUsdRate()
		} else if strings.Contains(msg.Text, "/eurate") {
			resp = fun.GetEuRate()
		} else if strings.Contains(msg.Text, "/about") {
			resp = "Типичный представитель расы киборгов, внедренных в человеческое общество. \nЛюбит пиво и ненавидит Пабло. \nАктивный и публичный гей, либерал, консерватор, традиционалист, социалист, экзгибиционист. \nПредпочитает небинарные трансгендерные отношения из-за биполярной настройки, а так же имеет латентную склонностью к натурализму."
		} else if strings.Contains(msg.Text, "/rate") {
			rand.Seed(time.Now().Unix())
			resp = "Ваш высер оценен на: " + strconv.Itoa(rand.Intn(11)) + "\nПоздравляю!"
			log.Println(resp)
		} else if strings.Contains(msg.Text, "/whathappennext") {
			resp = "В ближайшее время " + fun.RandomEventsGenerator()
		} else if strings.Contains(msg.Text, "/what2doinweekend") {
			resp = fun.RandomEventsGenerator4Weekend()
		} else {
			resp = "He знаю чё ты хочешь, иди на хуй!"
		}
		SendMessage(update.Message.Chat.ID, resp, bot)
	}
}
