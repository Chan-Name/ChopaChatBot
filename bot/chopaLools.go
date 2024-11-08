package bot

import (
	"fmt"
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) chopaMessageChecker(command []string, message *tgbotapi.Message) {

	if command[0] == "чопа" || command[0] == "Чопа" {
		chopaWord := b.chopaLools(command)
		msg := tgbotapi.NewMessage(message.Chat.ID, chopaWord)
		b.bot.Send(msg)
	}
}

func (b *Bot) chopaLools(command []string) string {

	switch command[1] {
	case "нюхай":
		return ChopaSniff(command)

	case "лизни", "лижи":
		return ChopaLick(command)

	case "соси":
		return ChopaSuck(command)

	case "ешь", "съешь":
		return ChopaEat(command)

	case "инфа":
		return ChopaInf()

	case "порно":
		return ChopaPorn()

	case "команды":
		return ChopaCommands()

	default:
		return "Чопа не знает такой команды"

	}
}

func ChopaSniff(command []string) string {
	return ChopaActions("нюхаю", command)
}

func ChopaLick(command []string) string {
	return ChopaActions("лижу", command)
}

func ChopaSuck(command []string) string {
	return ChopaActions("сосу", command)
}

func ChopaEat(command []string) string {
	return ChopaActions("ем", command)
}

func ChopaInf() string {
	return fmt.Sprintf("Инфа шанс %d%%", rand.Intn(101))
}

func ChopaPorn() string {
	return fmt.Sprintf("ПОРНО!\n%v", "https://www.xvideos.com/\n"+"https://rt.pornhub.com/\n"+
		"http://porno365.fun/\n"+"")
}

func ChopaCommands() string {
	return "Админ-команды: мут 10 минут(секунд, часов, лет), бан, разбан\nЮзер-команды(обязательна приписка чопа): нюхай, лизни, соси, ешь, инфа, порно, команды"

}

func ChopaActions(action string, command []string) string {
	toAction := command[2:]
	return fmt.Sprintf("*%s %s*", action, strings.Join(toAction, " "))
}
