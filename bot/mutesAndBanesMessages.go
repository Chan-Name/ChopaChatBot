package bot

import (
	"fmt"
)

func checkMessageForUseMute(command string, len int) bool {

	if len == 3 {
		if command == "мут" || command == "Мут" {
			return true
		}
	}
	return false
}

func (b *Bot) messageForMute(user string, chatID int64, date *MuteData) {
	msg := fmt.Sprintf("@%v %v, %d, %v", user,
		"is muted to", date.Time, date.Period)
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForBan(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"is banned")
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForUnBan(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"is unbanned now")
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForUnMute(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"is unmuted now")
	b.sendMessage(chatID, msg)
}
func (b *Bot) messageForNonAdmin(chatID int64) {
	msg := "You are not an admin, don't try"
	b.sendMessage(chatID, msg)
}
