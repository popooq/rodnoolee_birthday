package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/popooq/rodnoolee_birthday/internal/handler"
)

type tgBot struct {
	botAPI        *tgbotapi.BotAPI
	commands      map[string]func()
	handlerAdd    handler.Handler
	handlerGetAll handler.Handler
	handlerUpdate handler.Handler
}

func New(token string) *tgBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	tgbot := tgBot{
		botAPI: bot,
	}
}

func (bot *tgBot) Process() {
	bot.botAPI.Debug = true

	log.Printf("Authorized on account %s", &bot.botAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.botAPI.Send(msg)
		}
	}
}
