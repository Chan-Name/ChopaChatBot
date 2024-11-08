package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) adminChecker(userID int, chatID int64) bool {

	admins, err := b.bot.GetChatAdministrators(tgbotapi.ChatConfig{
		ChatID: chatID,
	})
	if err != nil {
		log.Println(err)
	}

	for _, admin := range admins {
		if userID == admin.User.ID {
			return true
		}
	}
	return false
}

func (b *Bot) AdminCommandsChecker(command []string, message *tgbotapi.Message) {
	if message.ReplyToMessage != nil {

		switch command[0] {
		case "мут", "Мут":
			b.mute(message.ReplyToMessage.From, message, command)

		case "бан", "Бан":
			b.ban(message.ReplyToMessage.From, message)

		case "разбан", "Разбан":
			b.unBan(message.ReplyToMessage.From, message)

		case "размут", "Размут":
			b.unMute(message.ReplyToMessage.From, message)
		}
	}
}

func (b *Bot) allCommandsChecker(command []string, message *tgbotapi.Message) {
	b.chopaMessageChecker(command, message)
	b.AdminCommandsChecker(command, message)
}
