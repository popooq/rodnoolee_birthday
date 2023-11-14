package allbirthdays

import (
	"github.com/popooq/rodnoolee_birthday/internal/repository"
)

type allBirthdays struct {
	repo repository.UserRepo
}

func New(repo repository.UserRepo) AllBirthdays {
	return &allBirthdays{
		repo: repo,
	}
}

func (a *allBirthdays) Handle() ([]repository.Birthday, error) {
	birthdays, err := a.repo.GetAllBirthdays()
	if err != nil {
		return nil, err
	}

	return birthdays, nil
}
