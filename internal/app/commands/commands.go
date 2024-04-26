package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muradrmagomedov/bot/internal/service/product"
)

type Commander struct {
	bot             *tgbotapi.BotAPI
	productServices *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productServices *product.Service) *Commander {
	return &Commander{
		bot:             bot,
		productServices: productServices,
	}
}

func (c *Commander) Help(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "/help - help\n"+"/list - product list")
	c.bot.Send(reply)
}

func (c *Commander) List(msg *tgbotapi.Message) {
	products := c.productServices.List()
	productsMsgText := "Here all your products\n\n"
	for _, product := range products {
		productsMsgText += product.Title + "\n"
	}
	reply := tgbotapi.NewMessage(msg.Chat.ID, productsMsgText)
	c.bot.Send(reply)
}

func (c *Commander) Default(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	reply.ReplyToMessageID = msg.MessageID
	c.bot.Send(reply)
}
