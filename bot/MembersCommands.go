package bot

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) membersCommandChecker(command []string, message *tgbotapi.Message) {
	switch strings.ToLower(command[0]) {
	case "админы":
		b.adminsOutput(message.Chat.ID)
	}
}

func (b *Bot) adminsOutput(chatID int64) {
	admins, err := b.bot.GetChatAdministrators(tgbotapi.ChatConfig{
		ChatID: chatID,
	})

	if err != nil {
		log.Println(err)
	}
	adminsToSend := "Админы:\n"
	for _, admin := range admins {
		adminsToSend += fmt.Sprintf("@%v\n", admin.User)
	}
	b.sendMessage(chatID, adminsToSend)
}
