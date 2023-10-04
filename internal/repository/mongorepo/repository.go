package mongorepo

import (
	"context"
	"log"

	"github.com/popooq/rodnoolee_birthday/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	client     mongo.Client
	address    string
	collection *mongo.Collection
	context    context.Context
}

func New(ctx context.Context, dba string) (repository.UserRepo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dba))
	if err != nil {
		return nil, err
	}

	collection := client.Database("bdays").Collection("rodnoolee")

	return &mongoRepo{
		client:     *client,
		address:    dba,
		collection: collection,
		context:    ctx,
	}, nil
}

func (m *mongoRepo) InsertBirthday(birthday repository.Birthday) error {
	_, err := m.collection.InsertOne(m.context, birthday)

	return err
}

func (m *mongoRepo) GetAllBirthdays() ([]repository.Birthday, error) {
	cursor, err := m.collection.Find(m.context, bson.D{})

	var birthdays []repository.Birthday

	for cursor.Next(m.context) {
		var birthday repository.Birthday
		err := cursor.Decode(&birthday)
		if err != nil {
			log.Fatal(err)
		}

		birthdays = append(birthdays, birthday)
	}
	return birthdays, err
}

func (m *mongoRepo) UpdateBirthday(birthday repository.Birthday) error {
	_, err := m.collection.UpdateOne(m.context, bson.D{
		{Key: "Rodnoolya", Value: birthday.Rodnoolya},
	}, birthday)

	return err
}
