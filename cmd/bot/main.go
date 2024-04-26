package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/muradrmagomedov/bot/internal/app/commands"
	"github.com/muradrmagomedov/bot/internal/service/product"
)

func main() {
	//Загружаем env файл
	godotenv.Load("../../.env")
	token := os.Getenv("TOKEN")

	productServices := product.NewService()

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

	commander := commands.NewCommander(bot, productServices)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			//обработка команды help
			switch update.Message.Command() {
			case "help":
				commander.Help(update.Message)
			case "list":
				commander.List(update.Message)
			default:
				commander.Default(update.Message)
			}
		}
	}
}
