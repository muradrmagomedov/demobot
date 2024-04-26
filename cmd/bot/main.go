package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/muradrmagomedov/bot/internal/service/product"
)

func main() {
	//Загружаем env файл
	godotenv.Load("../../.env")
	token := os.Getenv("TOKEN")

	service := product.NewService()

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
			case "list":
				listCommand(bot, update.Message, service)
			default:
				defaultAnswer(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "/help - help\n"+"/list - product list")
	bot.Send(reply)
}

func listCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, service *product.Service) {
	products := service.List()
	productsMsgText := "Here all your products\n\n"
	for _, product := range products {
		productsMsgText += product.Title + "\n"
	}
	reply := tgbotapi.NewMessage(msg.Chat.ID, productsMsgText)
	bot.Send(reply)
}

func defaultAnswer(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID
	bot.Send(reply)
}
