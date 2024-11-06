package bot

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b *Bot) Start() {

	updates, err := b.initStart()
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		message := update.Message
		if message != nil {

			command := strings.Fields(message.Text)

			if len(command) > 0 && command[0] != "" {
				switch command[0] {
				case "чопа", "Чопа":
					chopaWord := b.chopaLools(command)
					msg := tgbotapi.NewMessage(message.Chat.ID, chopaWord)
					b.bot.Send(msg)
				}

				if message.ReplyToMessage != nil && command[3] != "" {
					replyMessage := message.ReplyToMessage
					switch command[0] {
					case "бан", "Бан":
						timeForBan := timeForMuteOrBan(command[1:])

						b.ban(replyMessage.From, message.Chat, int64(timeForBan.Time))

						msg := tgbotapi.NewMessage(message.Chat.ID, messageForBan(replyMessage.From.UserName, &timeForBan))
						b.bot.Send(msg)

					case "мут", "Мут":
						timeForMute := timeForMuteOrBan(command[1:])
						b.mute(replyMessage.From, message.Chat, int64(timeForMute.Time))

						msg := tgbotapi.NewMessage(message.Chat.ID, messageForMute(replyMessage.From.UserName, &timeForMute))
						b.bot.Send(msg)
					}

				} else if message.ReplyToMessage != nil && command[3] != "" && command[1] == "бан" ||
					message.ReplyToMessage != nil && command[3] != "" && command[1] == "мут" {

					msg := tgbotapi.NewMessage(message.Chat.ID, "Неправильный вид даты и времени\n Пример: Мут(Бан) 10 минут")
					b.bot.Send(msg)
				}

			}
		}

	}
}

func (b *Bot) initStart() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}
