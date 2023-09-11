package handlers

import "github.com/popooq/rodnoolee_birthday/internal/types"

type storageRepo interface {
	GetAllBirthdays() ([]types.Birthday, error)
	InsertBirthday(birthday types.Birthday) error
}
