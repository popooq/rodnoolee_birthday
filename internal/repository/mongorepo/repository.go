package mongorepo

import (
	"context"
	"log"

	"github.com/popooq/rodnoolee_birthday/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	client     mongo.Client
	address    string
	collection *mongo.Collection
	context    context.Context
}

func New(ctx context.Context, dba string) (*MongoRepo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dba))
	if err != nil {
		return nil, err
	}

	collection := client.Database("bdays").Collection("rodnoolee")

	return &MongoRepo{
		client:     *client,
		address:    dba,
		collection: collection,
		context:    ctx,
	}, nil
}

func (m *MongoRepo) InsertBirthday(birthday types.Birthday) error {
	_, err := m.collection.InsertOne(m.context, birthday)

	return err
}

func (m *MongoRepo) GetAllBirthdays() ([]types.Birthday, error) {
	cursor, err := m.collection.Find(m.context, bson.D{})

	var birthdays []types.Birthday

	for cursor.Next(m.context) {
		var birthday types.Birthday
		err := cursor.Decode(&birthday)
		if err != nil {
			log.Fatal(err)
		}

		birthdays = append(birthdays, birthday)
	}
	return birthdays, err
}
