package bot

import (
	"fmt"
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) chopaMessageChecker(command []string, message *tgbotapi.Message) {

	if strings.ToLower(command[0]) == "чопа" {
		chopaWord := b.chopaLools(command)
		msg := tgbotapi.NewMessage(message.Chat.ID, chopaWord)
		b.bot.Send(msg)
	}
}

func (b *Bot) chopaLools(command []string) string {

	switch strings.ToLower(command[1]) {

	case "инфа":
		return ChopaInf()

	case "команды":
		return ChopaCommands()

	default:
		return "Чопа не знает такой команды"

	}
}

func ChopaInf() string {
	return fmt.Sprintf("Инфа шанс %d%%", rand.Intn(101))
}

func ChopaCommands() string {
	return "Админ-команды: мут 10 секунд(минут, часов, недель, месяцев, лет), бан, разбан\nЮзер-команды: админы\n Чопа-команды(обязательна приписка чопа): инфа, команды"

}

func ChopaActions(action string, command []string) string {
	toAction := command[2:]
	return fmt.Sprintf("*%s %s*", action, strings.Join(toAction, " "))
}
