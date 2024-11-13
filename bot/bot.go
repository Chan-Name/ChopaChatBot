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
				b.allCommandsChecker(command, message)
			}

		}
	}
}

func (b *Bot) initStart() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) sendMessage(ID int64, message string) {
	msg := tgbotapi.NewMessage(ID, message)
	b.bot.Send(msg)
}
