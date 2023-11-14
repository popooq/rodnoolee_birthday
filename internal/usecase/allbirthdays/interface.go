package allbirthdays

import "github.com/popooq/rodnoolee_birthday/internal/repository"

type AllBirthdays interface {
	Handle() ([]repository.Birthday, error)
}
