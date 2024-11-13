package bot

import (
	"log"
	"strings"

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

		switch strings.ToLower(command[0]) {
		case "повысить":
			b.promoteToAdmin(message.ReplyToMessage.From, message)

		case "понизить":
			b.lowerAdmin(message.ReplyToMessage.From, message)

		case "мут":
			b.mute(message.ReplyToMessage.From, message, command)

		case "бан":
			b.ban(message.ReplyToMessage.From, message)

		case "разбан":
			b.unBan(message.ReplyToMessage.From, message)

		case "размут":
			b.unMute(message.ReplyToMessage.From, message)
		}
	}
}

func (b *Bot) allCommandsChecker(command []string, message *tgbotapi.Message) {
	b.membersCommandChecker(command, message)
	b.chopaMessageChecker(command, message)
	b.AdminCommandsChecker(command, message)
}
