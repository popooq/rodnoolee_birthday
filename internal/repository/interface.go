package repository

type Birthday struct {
	ID        string `bson:"_id"`
	Rodnoolya string `bson:"rodnoolya"`
	Birthday  string `bson:"bday"`
}

type UserRepo interface {
	InsertBirthday(birthday Birthday) error
	GetAllBirthdays() ([]Birthday, error)
	UpdateBirthday(rodnoolya string, birthday Birthday) error
}
