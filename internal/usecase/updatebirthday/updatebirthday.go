package updatebirthday

import (
	"time"

	"github.com/popooq/rodnoolee_birthday/internal/domain"
	"github.com/popooq/rodnoolee_birthday/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const shortForm = time.DateOnly

type updateBirthday struct {
	repo repository.UserRepo
}

func New(repo repository.UserRepo) UpdateBirthday {
	return &updateBirthday{
		repo: repo,
	}
}
func (u *updateBirthday) Handle(message domain.TgMessage) error {
	dt, err := time.Parse(shortForm, message.MessageText)
	if err != nil {
		return err
	}
	mdt := primitive.NewDateTimeFromTime(dt)

	birthday := repository.Birthday{
		Rodnoolya: message.Username,
		Birthday:  mdt,
	}
	err = u.repo.UpdateBirthday(birthday)
	if err != nil {
		return err
	}

	return nil
}
