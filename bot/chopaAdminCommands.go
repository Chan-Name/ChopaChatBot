package bot

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type timeForMuteBan struct {
	Time   int
	Period string
}

func (b *Bot) ban(userForBan *tgbotapi.User, chat *tgbotapi.Chat, date int64) {

	memberConfig := tgbotapi.ChatMemberConfig{
		ChatID:          chat.ID,
		ChannelUsername: chat.UserName,
		UserID:          userForBan.ID,
	}

	kickConfig := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: memberConfig,
		UntilDate:        time.Now().Unix() + date,
	}

	_, err := b.bot.KickChatMember(kickConfig)
	if err != nil {
		log.Fatal(err)
	}

}

func (b *Bot) mute(userForMute *tgbotapi.User, chat *tgbotapi.Chat, date int64) {

	memberConfig := tgbotapi.ChatMemberConfig{
		ChatID:          chat.ID,
		ChannelUsername: chat.UserName,
		UserID:          userForMute.ID,
	}

	restrictConfig := tgbotapi.RestrictChatMemberConfig{
		ChatMemberConfig: memberConfig,
		UntilDate:        time.Now().Unix() + date,
	}

	_, err := b.bot.RestrictChatMember(restrictConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func messageForBan(userName string, date *timeForMuteBan) string {
	return fmt.Sprintf("@%v %v, %d, %v", userName,
		"is banned to", date.Time, date.Period)
}

func messageForMute(userName string, date *timeForMuteBan) string {
	return fmt.Sprintf("@%v %v, %d, %v", userName,
		"is muted to", date.Time, date.Period)
}

func timeForMuteOrBan(command []string) timeForMuteBan {

	startTime, err := strconv.Atoi(command[3])
	if err != nil {
		log.Fatal(err)
	}

	switch command[3] {
	case "минут", "минута", "минуты", "минуту":
		return timeForMuteBan{
			Time:   startTime * 60,
			Period: command[3],
		}

	case "час", "часа", "часов":
		return timeForMuteBan{
			Time:   startTime * 3600,
			Period: command[3],
		}

	case "день", "дня", "дней":
		return timeForMuteBan{
			Time:   startTime * 86400,
			Period: command[3],
		}

	case "недель", "недели", "неделя", "неделю":
		return timeForMuteBan{
			Time:   startTime * 604800,
			Period: command[3],
		}

	case "месяц", "месяца", "месяцов":
		return timeForMuteBan{
			Time:   (startTime * 86400) * 30,
			Period: command[3],
		}

	case "год", "года", "лет":
		return timeForMuteBan{
			Time:   (startTime * 86400) * 365,
			Period: command[3],
		}

		// default тут - заглушка, что никогда не сработает, чтобы компилятор не выдавал ошибку
	default:
		return timeForMuteBan{
			Time:   0,
			Period: "0",
		}
	}
}
