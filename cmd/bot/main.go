package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	//Загружаем env файл
	godotenv.Load("../../.env")
	token := os.Getenv("TOKEN")

	//создаем объект бота и подключаемся по токену
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	//включаем дебаг мод
	bot.Debug = true
	//сообщение об успешной авторизации
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			//обработка команды help
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			default:
				defaultAnswer(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "You need some help?")
	bot.Send(reply)
}

func defaultAnswer(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID
	bot.Send(reply)
}
