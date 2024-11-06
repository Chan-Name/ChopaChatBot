package main

import (
	"chopa/bot"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	Bot, err := tgbotapi.NewBotAPI(jsonTokenDecode())
	if err != nil {
		log.Fatal(err)
	}

	b := bot.NewBot(Bot)
	b.Start()
}
