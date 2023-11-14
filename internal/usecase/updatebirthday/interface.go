package updatebirthday

import "github.com/popooq/rodnoolee_birthday/internal/domain"

type UpdateBirthday interface {
	Handle(message domain.TgMessage) error
}
