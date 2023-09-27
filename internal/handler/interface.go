package handler

import "github.com/popooq/rodnoolee_birthday/internal/domain"

type Handler interface {
	Handle(command string, message domain.TgMessage) error
}
