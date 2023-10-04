package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

//go:generate mockgen -destination ./mocks/client.go -package=mocks -source interface.go -package=mockrepo

type Birthday struct {
	ID        string             `bson:"_id"`
	Rodnoolya string             `bson:"rodnoolya"`
	Birthday  primitive.DateTime `bson:"bday"`
}

type UserRepo interface {
	InsertBirthday(birthday Birthday) error
	GetAllBirthdays() ([]Birthday, error)
	UpdateBirthday(birthday Birthday) error
}
