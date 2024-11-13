package bot

import (
	"fmt"
	"strings"
)

func checkMessageForUseMute(command string, len int) bool {

	if len == 3 {
		if strings.ToLower(command) == "мут" {
			return true
		}
	}
	return false
}

func (b *Bot) messageForMute(user string, chatID int64, date *MuteData) {
	msg := fmt.Sprintf("@%v %v %d %v", user,
		"будет молчать", date.TimeForMessage, date.Period)
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForBan(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"забанен")
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForUnBan(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"разбанен")
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForUnMute(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"закончил молчать")
	b.sendMessage(chatID, msg)
}
func (b *Bot) messageForNonAdmin(chatID int64) {
	msg := "Команда доступна только администраторам"
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForMemberPromoteToAdmin(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"повышен до админа")
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForMemberPromoteAdminToAdmin(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"уже админ")
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForAdminLowerToMember(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"понижен до участника")
	b.sendMessage(chatID, msg)
}

func (b *Bot) messageForAdminLowerMemberToMember(user string, chatID int64) {
	msg := fmt.Sprintf("@%v %v", user,
		"уже участник")
	b.sendMessage(chatID, msg)
}
