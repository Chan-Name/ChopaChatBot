package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) promoteToAdmin(user *tgbotapi.User, message *tgbotapi.Message) {
	if !b.adminChecker(message.From.ID, message.Chat.ID) {
		b.messageForNonAdmin(message.Chat.ID)
		return
	}
	if !b.adminChecker(user.ID, message.Chat.ID) {
		b.messageForMemberPromoteAdminToAdmin(user.UserName, message.Chat.ID)
	} else {

		promote := promoteConfigCreator(user, message.Chat, true)

		_, err := b.bot.PromoteChatMember(promote)
		if err != nil {
			log.Println(err)
		}
		b.messageForMemberPromoteToAdmin(user.UserName, message.Chat.ID)
	}
}

func (b *Bot) lowerAdmin(user *tgbotapi.User, message *tgbotapi.Message) {
	if !b.adminChecker(message.From.ID, message.Chat.ID) {
		b.messageForNonAdmin(message.Chat.ID)
		return
	}
	if !b.adminChecker(user.ID, message.Chat.ID) {
		b.messageForAdminLowerMemberToMember(user.UserName, message.Chat.ID)
	} else {

		promote := promoteConfigCreator(user, message.Chat, false)
		_, err := b.bot.PromoteChatMember(promote)
		if err != nil {
			log.Println(err)
		}
		b.messageForAdminLowerToMember(user.UserName, message.Chat.ID)
	}
}
