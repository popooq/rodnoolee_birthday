package mappers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/popooq/rodnoolee_birthday/internal/domain"
)

func TgMessageToDomain(tgMessage tgbotapi.Message) domain.TgMessage {

	return domain.TgMessage{
		MessageText: tgMessage.Text,
	}
}
