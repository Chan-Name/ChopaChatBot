package bot

import (
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MuteData struct {
	Time           int
	TimeForMessage int
	Period         string
}

func (b *Bot) ban(userToBan *tgbotapi.User, message *tgbotapi.Message) {

	if b.adminChecker(message.From.ID, message.Chat.ID) {

		kickConfig := kickConfigCreator(userToBan, message.Chat)

		_, err := b.bot.KickChatMember(kickConfig)
		if err != nil {
			log.Println(err)
		}
		b.messageForBan(userToBan.UserName, message.Chat.ID)
	} else {
		b.messageForNonAdmin(message.Chat.ID)
	}

}

func (b *Bot) unBan(userToUnBan *tgbotapi.User, message *tgbotapi.Message) {

	if b.adminChecker(message.From.ID, message.Chat.ID) {

		kickConfig := kickConfigCreator(userToUnBan, message.Chat)

		_, err := b.bot.KickChatMember(kickConfig)
		if err != nil {
			log.Println(err)
		} else {
			b.messageForUnBan(userToUnBan.UserName, message.Chat.ID)
		}
	} else {
		b.messageForNonAdmin(message.Chat.ID)
	}
}

func (b *Bot) mute(user *tgbotapi.User, message *tgbotapi.Message, command []string) {
	if b.adminChecker(message.From.ID, message.Chat.ID) {

		if checkMessageForUseMute(command[0], len(command)) {

			date := timeForMute(command[1:])

			restrictConfigToMute := restrictConfigToMuteCreator(user, message.Chat)
			restrictConfigToMute.UntilDate += int64(date.Time)

			_, err := b.bot.RestrictChatMember(restrictConfigToMute)
			if err != nil {
				log.Println(err)
			}

			go func() {

				for i := 0; i < date.Time; {
					time.Sleep(1 * time.Second)
					i += 1
				}
				b.unMute(user, message)
			}()

			b.messageForMute(user.UserName, message.Chat.ID, date)

		} else {
			b.sendMessage(message.Chat.ID, "Неправильный формат даты и времени\nПример: Мут 10 минут")
		}

	}
}

func (b *Bot) unMute(user *tgbotapi.User, message *tgbotapi.Message) {

	if b.adminChecker(message.From.ID, message.Chat.ID) {

		restrictConfigToUnMute := restrictConfigToUnMuteCreator(user, message.Chat)

		_, err := b.bot.RestrictChatMember(restrictConfigToUnMute)
		if err != nil {
			log.Println(err)
		}
		b.messageForUnMute(user.UserName, message.Chat.ID)
	} else {
		b.messageForNonAdmin(message.Chat.ID)
	}
}

func timeForMute(command []string) *MuteData {

	startTime, err := strconv.Atoi(command[0])
	if err != nil {
		log.Println(err)
	}

	switch strings.ToLower(command[1]) {
	case "секунд", "секунда", "секунду", "секунды":
		return &MuteData{
			Time:           startTime,
			TimeForMessage: startTime,
			Period:         command[1],
		}
	case "минут", "минута", "минуты", "минуту":
		return &MuteData{
			Time:           startTime * 60,
			TimeForMessage: startTime,
			Period:         command[1],
		}

	case "час", "часа", "часов":
		return &MuteData{
			Time:           startTime * 3600,
			TimeForMessage: startTime,
			Period:         command[1],
		}

	case "день", "дня", "дней":
		return &MuteData{
			Time:           startTime * 86400,
			TimeForMessage: startTime,
			Period:         command[1],
		}

	case "недель", "недели", "неделя", "неделю":
		return &MuteData{
			Time:           startTime * 604800,
			TimeForMessage: startTime,
			Period:         command[1],
		}

	case "месяц", "месяца", "месяцов":
		return &MuteData{
			Time:           (startTime * 86400) * 30,
			TimeForMessage: startTime,
			Period:         command[1],
		}

	case "год", "года", "лет":
		return &MuteData{
			Time:           (startTime * 86400) * 365,
			TimeForMessage: startTime,
			Period:         command[1],
		}

	default:
		return nil
	}
}
