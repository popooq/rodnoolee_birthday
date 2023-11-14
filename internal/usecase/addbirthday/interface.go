package addbirthday

import "github.com/popooq/rodnoolee_birthday/internal/domain"

type AddBirthday interface {
	Handle(message domain.TgMessage) error
}
